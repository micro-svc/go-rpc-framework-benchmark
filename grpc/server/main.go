package main

import (
	"context"
	"flag"
	"net"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/grpc/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/phuslu/log"
	"google.golang.org/grpc"
)

var (
	// flags
	address = flag.String("addr", "127.0.0.1:5678", "server listen address")
)

type Server struct {
	model.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, req *model.Message) (*model.Message, error) {
	rsp := &model.Message{}
	*rsp = *req
	return rsp, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		ulog.FatalIfError(err)
	}

	srv := grpc.NewServer()
	model.RegisterHelloServer(srv, new(Server))
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(srv.Serve(lis))
}
