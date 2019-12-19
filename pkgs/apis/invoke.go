package apis

import (
	"github.com/ckeyer/diego/pkgs/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GinInvoke(ctx *gin.Context, function interface{}) {
	if err := app.Invoke(function); err != nil {
		logrus.Errorf("%s", err)
		InternalServerErr(ctx, err)
		return
	}
}
