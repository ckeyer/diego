package api

import (
	"github.com/ckeyer/diego/pkgs/projects"
	"github.com/gin-gonic/gin"
)

// usersRouters ...
func usersRouters(gr *gin.RouterGroup) {
	gr.GET("projects", projects.ListProjects)
}
