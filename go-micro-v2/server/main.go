package main

import (
	"context"
	"flag"

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
	"github.com/phuslu/log"
)

var (
	// flags
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "grpc", "server transport") // grpc, http

	newServer = server.NewServer
	newClient = client.NewClient
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *model.Message, rsp *model.Message) error {
	*rsp = *req
	return nil
}

func main() {
	flag.Parse()

	switch *transport {
	case "grpc":
		newServer = srvgrpc.NewServer
		newClient = cligrpc.NewClient
	case "http":
		ts.DefaultTransport = ts.NewTransport()
	default:
		log.Fatal().Msg("flag transport not support")
	}

	server.DefaultServer = newServer(
		server.Codec("application/json", json.NewCodec),
		server.Codec("application/protobuf", proto.NewCodec),
	)
	client.DefaultClient = newClient(
		client.Codec("application/json", json.NewCodec),
		client.Codec("application/protobuf", proto.NewCodec),
	)

	svc := micro.NewService(
		micro.Name("hello"),
		micro.Address(*address),
	)

	ulog.FatalIfError(model.RegisterHelloHandler(svc.Server(), new(Server)))
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(svc.Run())
}
