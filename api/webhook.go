package api

import (
	"errors"

	"github.com/ckeyer/diego/hacks/webhook"
	"github.com/ckeyer/diego/pkgs/apis"
	"github.com/gin-gonic/gin"
)

// DoWebhook .
func DoWebhook(ctx *gin.Context) {
	param := ctx.Param("cmd")
	if param == "" {
		apis.InternalServerErr(ctx, errors.New("empty cmd"))
		return
	}
	webhook.Exec(param)
}
