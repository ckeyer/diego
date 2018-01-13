package api

import (
	"strings"
	"time"

	"github.com/ckeyer/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// MDCors middleware for CORS.
func MDCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Limit", "Offset", "Origin", "Accept", "X-Signature"},
		ExposeHeaders:    []string{"Content-Length", "Accept-Encoding"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}

// MDLogger middleware for http logger.
func MDLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		logent := logrus.WithFields(logrus.Fields{
			"Method": ctx.Request.Method,
			"URL":    ctx.Request.URL.Path,
			"Remote": ctx.Request.RemoteAddr,
			"Status": ctx.Writer.Status(),
		})

		for _, prefix := range []string{API_PREFIX, UI_PREFIX} {
			if strings.HasPrefix(ctx.Request.URL.Path, prefix) {
				logent.Infof("%.6f", time.Now().Sub(start).Seconds())
				return
			}
		}
		logent.Debugf("%.6f", time.Now().Sub(start).Seconds())
	}
}

func MDRecovery() gin.HandlerFunc {
	return gin.Recovery()
}
