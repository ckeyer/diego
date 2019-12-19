package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/diego/version"

	"github.com/ckeyer/diego/pkgs/apis/ginmd"
	"github.com/gin-gonic/gin"
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
	gr.GET("/_ping", getVersion)
	usersRouters(gr.Group(""))
}

// webhook api router
func webhookRoute(gr *gin.RouterGroup) {
	gr.POST("/webhook/:cmd", DoWebhook())
}

func getVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, version.Map())
}

func decodeBody(ctx *gin.Context, v interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(v)
}
