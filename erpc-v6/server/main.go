package main

import (
	"flag"
	"strings"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	"github.com/gxxgle/go-utils/conver"
	ulog "github.com/gxxgle/go-utils/log"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/codec"
	"github.com/phuslu/log"
)

var (
	// flags
	address   = flag.String("addr", "127.0.0.1:5678", "server listen address")
	transport = flag.String("transport", "tcp", "server transport") // tcp
	xcodec    = flag.String("codec", "json", "server codec")        // json, protobuf

	cfgCodec string
)

type Server struct {
	erpc.CallCtx
}

func (s *Server) Hello(req *model.Message) (*model.Message, *erpc.Status) {
	rsp := &model.Message{}
	*rsp = *req
	return rsp, nil
}

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

	addres := strings.Split(*address, ":")
	if len(addres) != 2 {
		log.Fatal().Msg("flag addr not support")
	}

	srv := erpc.NewPeer(erpc.PeerConfig{
		Network:          *transport,
		LocalIP:          addres[0],
		ListenPort:       uint16(conver.IntMust(addres[1])),
		DefaultBodyCodec: cfgCodec,
	})

	srv.RouteCall(new(Server))
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(srv.ListenAndServe())
}
