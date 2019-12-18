package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundErr(ctx *gin.Context, err interface{}) {
	abortErr(ctx, http.StatusNotFound, err)
}

func BadRequestErr(ctx *gin.Context, err interface{}) {
	abortErr(ctx, http.StatusBadRequest, err)
}

func InternalServerErr(ctx *gin.Context, err interface{}) {
	abortErr(ctx, http.StatusInternalServerError, err)
}

func abortErr(ctx *gin.Context, code int, err interface{}) {
	ctx.AbortWithStatusJSON(code, map[string]interface{}{"error": err})
}
