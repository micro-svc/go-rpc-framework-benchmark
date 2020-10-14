package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
)

var (
	// flags
	clients  = flag.Int("clients", 100, "concurrency client amount")
	requests = flag.Int("requests", 1000, "request amount per client")
	address  = flag.String("addr", "127.0.0.1:5678", "server address")
)

func main() {
	flag.Parse()

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

func newClients(amount int) []*http.Client {
	clis := make([]*http.Client, amount)
	for i := 0; i < amount; i++ {
		clis[i] = newHTTPClient()
		call(clis[i]) // warmup
	}
	return clis
}

func call(cli *http.Client) {
	bs, _ := json.Marshal(model.Example)
	httprsp, err := cli.Post("http://"+*address, "application/json", bytes.NewBuffer(bs))
	ulog.FatalIfError(err)
	defer httprsp.Body.Close()
	body, _ := ioutil.ReadAll(httprsp.Body)
	rsp := &model.Message{}
	ulog.FatalIfError(json.Unmarshal(body, rsp))
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
