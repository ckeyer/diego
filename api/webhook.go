package api

import (
	"errors"

	"github.com/ckeyer/diego/tools/webhook"
	"github.com/gin-gonic/gin"
)

func DoWebhook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("cmd")
		if param == "" {
			InternalServerErr(ctx, errors.New("empty cmd"))
			return
		}
		webhook.Exec(param)
	}
}
