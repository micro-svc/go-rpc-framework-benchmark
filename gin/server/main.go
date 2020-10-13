package main

import (
	"flag"
	"net/http"

	"github.com/micro-svc/go-rpc-framework-benchmark/model/standard/model"

	"github.com/gin-gonic/gin"
	ulog "github.com/gxxgle/go-utils/log"
	"github.com/phuslu/log"
)

var (
	// flags
	address = flag.String("addr", "127.0.0.1:5678", "server listen address")
)

type Server struct{}

func (s *Server) Hello(ctx *gin.Context) {
	msg := model.Message{}
	ulog.FatalIfError(ctx.BindJSON(&msg))
	ctx.JSON(http.StatusOK, msg)
}

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)

	srv := &Server{}
	app := gin.New()
	app.POST("/", srv.Hello)
	log.Info().Str("address", *address).Msg("server running")
	ulog.FatalIfError(app.Run(*address))
}
