package api

import (
	"github.com/ckeyer/diego/pkgs/users"
	"github.com/gin-gonic/gin"
)

// usersRouters ...
func usersRouters(gr *gin.RouterGroup) {
	gr.GET("user", users.CreateUser)
}
