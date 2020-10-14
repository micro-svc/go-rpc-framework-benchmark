package main

import (
	"flag"
	"net"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/arpc/codec"
	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/lesismal/arpc"
	alog "github.com/lesismal/arpc/log"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
)

var (
	// flags
	clients   = flag.Int("clients", 100, "concurrency client amount")
	requests  = flag.Int("requests", 1000, "request amount per client")
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "tcp", "server transport") // tcp
	xcodec    = flag.String("codec", "json", "server codec")        // json, jsoniter, protobuf

	acodec codec.Codec
)

func main() {
	flag.Parse()
	alog.SetLogLevel(alog.LogLevelNone)

	switch *transport {
	case "tcp":
	default:
		log.Fatal().Msg("flag transport not support")
	}

	switch *xcodec {
	case "json":
		acodec = codec.Json
	case "jsoniter":
		acodec = codec.Jsoniter
	case "protobuf":
		acodec = codec.Protobuf
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

func newClients(amount int) []*arpc.Client {
	clis := make([]*arpc.Client, amount)
	for i := 0; i < amount; i++ {
		cli, err := arpc.NewClient(func() (net.Conn, error) {
			return net.DialTimeout(*transport, *address, time.Second*3)
		})
		ulog.FatalIfError(err)
		cli.Codec = acodec
		cli.Run()
		clis[i] = cli
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli *arpc.Client) {
	rsp := &model.Message{}
	err := cli.Call("Hello", model.Example, rsp, time.Second*30)
	ulog.FatalIfError(err)
	if !model.CheckExample(rsp) {
		log.Fatal().Str("response", rsp.String()).Msg("response not match")
	}
}
