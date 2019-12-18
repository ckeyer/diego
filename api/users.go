package api

import "github.com/gin-gonic/gin"

// usersRouters ...
func usersRouters(gr *gin.RouterGroup) {
	gr.GET("/")
}
