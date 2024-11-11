package api

import (
	"github.com/ckeyer/diego/pkgs/users"
	"github.com/gin-gonic/gin"
)

func usersRouters(gr *gin.RouterGroup) {
	gr.GET("", users.ListUsers)
	gr.POST("", users.CreateUser)
	gr.GET(":user_id", users.GetUserProfile)
}
