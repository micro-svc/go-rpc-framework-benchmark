package main

import (
	"context"
	"flag"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"
	"github.com/micro-svc/go-rpc-framework-benchmark/rpcx-v5/tls"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/phuslu/log"
	xlog "github.com/smallnest/rpcx/v5/log"
	"github.com/smallnest/rpcx/v5/server"
)

var (
	// flags
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "tcp", "server transport") // tcp, quic, utp

	options []server.OptionFn
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *model.Message, rsp *model.Message) error {
	*rsp = *req
	return nil
}

func main() {
	flag.Parse()
	xlog.SetDummyLogger()
	ulog.ColorConsole()

	switch *transport {
	case "tcp":
	case "quic", "utp":
		cfg, err := tls.NewConfig()
		ulog.FatalIfError(err)
		options = append(options, server.WithTLSConfig(cfg))
	default:
		log.Fatal().Msg("flag transport not support")
	}

	// server.UsePool = true
	server := server.NewServer(options...)
	ulog.FatalIfError(server.RegisterName("hello", new(Server), ""))
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(server.Serve(*transport, *address))
}
