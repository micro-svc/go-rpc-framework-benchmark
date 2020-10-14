package main

import (
	"flag"
	"net"

	"github.com/micro-svc/go-rpc-framework-benchmark/arpc/codec"
	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	ulog "github.com/gxxgle/go-utils/log"
	"github.com/lesismal/arpc"
	alog "github.com/lesismal/arpc/log"
	"github.com/phuslu/log"
)

var (
	// flags
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "tcp", "server transport") // tcp
	xcodec    = flag.String("codec", "json", "server codec")        // json, jsoniter, protobuf

	acodec codec.Codec
)

type Server struct{}

func (s *Server) Hello(ctx *arpc.Context) {
	req := &model.Message{}
	ulog.FatalIfError(ctx.Bind(req))
	ulog.FatalIfError(ctx.Write(req))
}

func main() {
	flag.Parse()
	alog.SetLogLevel(alog.LogLevelNone)

	switch *transport {
	case "tcp":
	default:
		log.Fatal().Msg("flag transport not support")
	}

	switch *xcodec {
	case "json":
		acodec = codec.Json
	case "jsoniter":
		acodec = codec.Jsoniter
	case "protobuf":
		acodec = codec.Protobuf
	default:
		log.Fatal().Msg("flag codec not support")
	}

	lis, err := net.Listen(*transport, *address)
	ulog.FatalIfError(err)

	srv := arpc.NewServer()
	srv.Codec = acodec
	srv.Handler.Handle("Hello", new(Server).Hello)
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(srv.Serve(lis))
}
