package main

import (
	"context"
	"crypto/tls"
	"flag"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"
	"github.com/micro-svc/go-rpc-framework-benchmark/rpcx-v5/codec"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
	"github.com/smallnest/rpcx/v5/client"
	xlog "github.com/smallnest/rpcx/v5/log"
	"github.com/smallnest/rpcx/v5/protocol"
)

var (
	// flags
	clients   = flag.Int("clients", 100, "concurrency client amount")
	requests  = flag.Int("requests", 1000, "request amount per client")
	address   = flag.String("addr", "127.0.0.1:5678", "server address")
	transport = flag.String("transport", "tcp", "server transport") // tcp, quic, utp
	xcodec    = flag.String("codec", "json", "server codec")        // json, protobuf, msgpack, jsoniter

	option = client.DefaultOption
)

func main() {
	flag.Parse()
	xlog.SetDummyLogger()
	codec.Register()

	switch *transport {
	case "tcp":
	case "quic", "utp":
		option.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	default:
		log.Fatal().Msg("flag transport not support")
	}

	switch *xcodec {
	case "json":
		option.SerializeType = protocol.JSON
	case "protobuf":
		option.SerializeType = protocol.ProtoBuffer
	case "msgpack":
		option.SerializeType = protocol.MsgPack
	case "jsoniter":
		option.SerializeType = codec.JSONiter
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

func newClients(amount int) []client.XClient {
	clis := make([]client.XClient, amount)
	discovery := client.NewPeer2PeerDiscovery(*transport+"@"+*address, "")
	for i := 0; i < amount; i++ {
		clis[i] = client.NewXClient("hello", client.Failfast, client.RoundRobin, discovery, option)
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli client.XClient) {
	rsp := &model.Message{}
	err := cli.Call(context.Background(), "Hello", model.Example, rsp)
	ulog.FatalIfError(err)
	if !model.CheckExample(rsp) {
		log.Fatal().Str("response", rsp.String()).Msg("response not match")
	}
}
