package api

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ckeyer/diego/api/view"
	"github.com/ckeyer/diego/storage"
	"github.com/ckeyer/logrus"
	"github.com/gin-gonic/gin"
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
	gs.Use(MDCors())
	gs.NoRoute(view.UI())
	gs.Use(MDRecovery(), MDLogger())

	apiRoute(gs.Group(API_PREFIX))

	dlRoute(gs.Group(API_PREFIX))

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

	gr.GET("/users", ListUsers())
	gr.POST("/users", CreateUser())
	gr.GET("/users/:name", GetUserProfile())
	gr.GET("/users/:name/check", CheckNamespace())

	gr.GET("/orgs", ListOrgs())
	gr.POST("/orgs", CreateOrg())
	gr.GET("/orgs/:name", GetOrgProfile())
	gr.GET("/orgs/:name/check", CheckNamespace())

	gr.GET("/projects/:namespace", ListProjects())
	gr.GET("/projects/:namespace/:name", GetProjectProfile())
	gr.POST("/projects/:namespace", CreateProject())
	gr.DELETE("/projects/:namespace/:name", RemoveProject())

	gr.POST("/files/:namespace/:name", UploadFile())
}

// dlRoute api router.
func dlRoute(gr *gin.RouterGroup) {
	// gr.GET("/:ns/:ns_name", todo)
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
	InternalServerErr(ctx, "todo")
}

func decodeBody(ctx *gin.Context, v interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(v)
}
