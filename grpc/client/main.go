package main

import (
	"context"
	"flag"
	"sync"
	"time"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/grpc/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/montanaflynn/stats"
	"github.com/phuslu/log"
	"google.golang.org/grpc"
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

func newClients(amount int) []model.HelloClient {
	clis := make([]model.HelloClient, amount)
	for i := 0; i < amount; i++ {
		conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			ulog.FatalIfError(err)
		}

		clis[i] = model.NewHelloClient(conn)
		call(clis[i]) // warmup
	}

	return clis
}

func call(cli model.HelloClient) {
	_, err := cli.Hello(context.Background(), model.Example)
	ulog.FatalIfError(err)
}
