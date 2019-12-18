package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/diego/pkgs/apis/ginmd"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	API_PREFIX = "/api"
	UI_PREFIX  = "/release"
)

// Serve start http server.
func Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	gs := gin.New()
	gs.Use(ginmd.MDCors())
	gs.Use(ginmd.MDRecovery(), ginmd.MDLogger())

	apiRoute(gs.Group(API_PREFIX))
	webhookRoute(gs.Group(API_PREFIX))

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

// webhook api router
func webhookRoute(gr *gin.RouterGroup) {
	gr.POST("/webhook/:cmd", DoWebhook())
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
