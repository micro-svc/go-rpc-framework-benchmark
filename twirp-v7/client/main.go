package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/twirp-v7/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
	"github.com/twitchtv/twirp"
)

var (
	// flags
	clients  = flag.Int("clients", 100, "concurrency client amount")
	requests = flag.Int("requests", 1000, "request amount per client")
	address  = flag.String("addr", "127.0.0.1:5678", "server address")
	codec    = flag.String("codec", "json", "server codec") // json, protobuf

	newClient func(string, model.HTTPClient, ...twirp.ClientOption) model.Hello
)

func main() {
	flag.Parse()

	switch *codec {
	case "json":
		newClient = model.NewHelloJSONClient
	case "protobuf":
		newClient = model.NewHelloProtobufClient
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

func newClients(amount int) []model.Hello {
	clis := make([]model.Hello, amount)
	for i := 0; i < amount; i++ {
		clis[i] = newClient("http://"+*address, newHTTPClient())
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli model.Hello) {
	rsp, err := cli.Hello(context.Background(), model.Example)
	ulog.FatalIfError(err)
	if !model.CheckExample(rsp) {
		log.Fatal().Str("response", rsp.String()).Msg("response not match")
	}
}

func newHTTPClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Jar:       jar,
		Transport: transport,
		Timeout:   2 * time.Minute,
	}
}
