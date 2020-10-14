package main

import (
	"flag"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	_ "github.com/gxxgle/go-utils/log"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/codec"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
)

var (
	// flags
	clients   = flag.Int("clients", 100, "concurrency client amount")
	requests  = flag.Int("requests", 1000, "request amount per client")
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "tcp", "server transport [tcp]")
	xcodec    = flag.String("codec", "json", "server codec [json, protobuf]")

	cfgCodec string
)

func main() {
	flag.Parse()
	erpc.SetLoggerLevel("OFF")

	switch *transport {
	case "tcp":
	default:
		log.Fatal().Msg("flag transport not support")
	}

	switch *xcodec {
	case "json":
		cfgCodec = new(codec.JSONCodec).Name()
	case "protobuf":
		cfgCodec = new(codec.ProtoCodec).Name()
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

func newClients(amount int) []erpc.Session {
	clis := make([]erpc.Session, amount)
	for i := 0; i < amount; i++ {
		peer := erpc.NewPeer(erpc.PeerConfig{
			Network:          *transport,
			DefaultBodyCodec: cfgCodec,
		})
		cli, stat := peer.Dial(*address)
		if !stat.OK() {
			log.Fatal().Msg(stat.String())
		}
		clis[i] = cli
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli erpc.Session) {
	rsp := &model.Message{}
	stat := cli.Call("/server/hello", model.Example, rsp).Status()
	if !stat.OK() {
		log.Fatal().Msg(stat.String())
	}
	if !model.CheckExample(rsp) {
		log.Fatal().Str("response", rsp.String()).Msg("response not match")
	}
}
