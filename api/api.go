package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/diego/storage"
	"github.com/ckeyer/logrus"
	"github.com/gin-gonic/gin"
	"gitlab.com/funxdata/commons/pkgs/ginmd"
)

var (
	stogr storage.Storager
)

const (
	API_PREFIX = "/api"
	UI_PREFIX  = "/release"
)

// Serve start http server.
func Serve(addr string, str storage.Storager) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	stogr = str

	gs := gin.New()
	gs.Use(ginmd.MDCors())
	gs.Use(ginmd.MDRecovery(), ginmd.MDLogger())

	apiRoute(gs.Group(API_PREFIX))

	err = http.Serve(lis, gs)
	if err != nil {
		return err
	}

	return nil
}

// apiRoute api router.
func apiRoute(gr *gin.RouterGroup) {
	gr.GET("/_ping", todo)

}

func todo(ctx *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"method": ctx.Request.Method,
		"path":   ctx.Request.URL.String(),
	}).Infof("ok.")

}

func decodeBody(ctx *gin.Context, v interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(v)
}
