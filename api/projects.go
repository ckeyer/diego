package api

import (
	"github.com/ckeyer/diego/pkgs/projects"
	"github.com/gin-gonic/gin"
)

// projectsRouters ...
func projectsRouters(gr *gin.RouterGroup) {
	gr.GET("", projects.ListProjects)
}
