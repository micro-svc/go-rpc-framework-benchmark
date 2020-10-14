package main

import (
	"context"
	"flag"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/micro-v2/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	cligrpc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/codec/json"
	"github.com/micro/go-micro/v2/codec/proto"
	"github.com/micro/go-micro/v2/server"
	srvgrpc "github.com/micro/go-micro/v2/server/grpc"
	ts "github.com/micro/go-micro/v2/transport"
	tsgrpc "github.com/micro/go-micro/v2/transport/grpc"
	tshttp "github.com/micro/go-micro/v2/transport/http"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
)

var (
	// flags
	clients   = flag.Int("clients", 100, "concurrency client amount")
	requests  = flag.Int("requests", 1000, "request amount per client")
	transport = flag.String("transport", "grpc", "server transport [grpc, http]")
	codec     = flag.String("codec", "json", "server codec [json, protobuf]")

	newServer = server.NewServer
	newClient = client.NewClient
)

func main() {
	flag.Parse()

	switch *transport {
	case "grpc":
		newServer = srvgrpc.NewServer
		newClient = cligrpc.NewClient
		ts.DefaultTransport = tsgrpc.NewTransport()
	case "http":
		ts.DefaultTransport = tshttp.NewTransport()
	default:
		log.Fatal().Msg("flag transport not support")
	}

	switch *codec {
	case "protobuf":
		server.DefaultServer = newServer(
			server.Codec("application/protobuf", proto.NewCodec),
		)
		client.DefaultClient = newClient(
			client.Codec("application/protobuf", proto.NewCodec),
			client.ContentType("application/protobuf"),
		)
	case "json":
		server.DefaultServer = newServer(
			server.Codec("application/json", json.NewCodec),
		)
		client.DefaultClient = newClient(
			client.Codec("application/json", json.NewCodec),
			client.ContentType("application/json"),
		)
	default:
		log.Fatal().Msg("flag codec not support")
	}

	var (
		wg        sync.WaitGroup
		total     = *clients * *requests
		clis      = newClients(*clients)
		durations = make([]time.Duration, total)
	)

	log.Info().Int("clients", *clients).Int("requests", *requests).Int("total", total).Msg("")

	start := time.Now()
	wg.Add(total)
	for i := range clis {
		go func(i int) {
			cli := clis[i]
			for j := 0; j < *requests; j++ {
				start := time.Now()
				call(cli)
				durations[i**requests+j] = time.Since(start)
				wg.Done()
			}
		}(i)
	}
	wg.Wait()

	tps := int64(time.Duration(total) * time.Second / time.Since(start))
	s := stats.LoadRawData(durations)
	min, _ := s.Min()
	max, _ := s.Max()
	mean, _ := s.Mean()
	median, _ := s.Median()
	log.Info().
		Int64("tps", tps).
		Dur("min", time.Duration(min)).
		Dur("max", time.Duration(max)).
		Dur("mean", time.Duration(mean)).
		Dur("median", time.Duration(median)).
		Msg("")
}

func newClients(amount int) []model.HelloService {
	clis := make([]model.HelloService, amount)
	for i := 0; i < amount; i++ {
		service := micro.NewService()
		mcli := service.Client()
		ulog.FatalIfError(mcli.Init(client.RequestTimeout(time.Second * 30)))
		clis[i] = model.NewHelloService("hello", mcli)
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli model.HelloService) {
	rsp, err := cli.Hello(context.Background(), model.Example)
	ulog.FatalIfError(err)
	if !model.CheckExample(rsp) {
		log.Fatal().Str("response", rsp.String()).Msg("response not match")
	}
}
