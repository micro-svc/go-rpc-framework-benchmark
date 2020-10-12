package main

import (
	"context"
	"flag"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/micro-v2/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	cligrpc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/server"
	srvgrpc "github.com/micro/go-micro/v2/server/grpc"
	"github.com/phuslu/log"
)

var (
	// flags
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "grpc", "server transport") // grpc
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *model.Message, rsp *model.Message) error {
	*rsp = *req
	return nil
}

func main() {
	flag.Parse()
	ulog.ColorConsole()

	switch *transport {
	case "grpc":
		server.DefaultServer = srvgrpc.NewServer()
		client.DefaultClient = cligrpc.NewClient()
	default:
		log.Fatal().Msg("flag transport not support")
	}

	service := micro.NewService(
		micro.Name("hello"),
		micro.Address(*address),
	)

	ulog.FatalIfError(model.RegisterHelloHandler(service.Server(), new(Server)))
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(service.Run())
}
