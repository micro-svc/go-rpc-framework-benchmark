package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/twirp-v7/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/phuslu/log"
)

var (
	// flags
	address = flag.String("addr", "127.0.0.1:5678", "server listen address")
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *model.Message) (*model.Message, error) {
	rsp := &model.Message{}
	*rsp = *req
	return rsp, nil
}

func main() {
	flag.Parse()
	ulog.ColorConsole()

	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(http.ListenAndServe(*address, model.NewHelloServer(new(Server))))
}
