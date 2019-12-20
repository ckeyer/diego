package api

import (
	"github.com/ckeyer/diego/pkgs/users"
	"github.com/gin-gonic/gin"
)

// projectsRouters ...
func projectsRouters(gr *gin.RouterGroup) {
	gr.GET("users", users.CreateUser)
}
