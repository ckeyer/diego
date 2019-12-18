package view

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	Index, _ = Asset("index.html")
)

func UI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := strings.TrimPrefix(ctx.Request.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		var (
			code        = http.StatusOK
			contentType = "text/html"
		)
		// ctx.Writer.Header().Add("Content-Encoding", "gzip")

		switch {
		case strings.HasSuffix(path, ".html"):
			contentType = "text/html"
		case strings.HasSuffix(path, ".js"):
			contentType = "application/x-javascript"
		case strings.HasSuffix(path, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(path, ".svg"):
			contentType = "text/xml"
		case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jepg"):
			contentType = "image/jpeg"
		}

		body, err := Asset(path)
		if err != nil {
			logrus.Debugf("get path %s failed, use index.", path)
			body = Index
		} else {
			logrus.Debugf("get path %s and return.", path)
		}

		ctx.Data(code, contentType, body)
	}
}
