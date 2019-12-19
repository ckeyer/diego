package users

import (
	"net/http"
	"time"

	"github.com/ckeyer/diego/pkgs/apis"
	"github.com/ckeyer/diego/pkgs/apis/validate"
	"github.com/ckeyer/diego/pkgs/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Joined  time.Time `json:"joined"`
	Updated time.Time `json:"updated"`
}

func (User) TableName() string {
	return "users"
}

type ListUserOption struct {
	apis.ListOption
}

func CreateUser(ctx *gin.Context) {
	err := app.Invoke(func(db *gorm.DB) {
		logrus.Info("createUser")
		user := &User{}
		if err := ctx.ShouldBindJSON(user); err != nil {
			apis.BadRequestErr(ctx, err.Error())
			return
		}
		if err := validate.IsDNS1035Label(user.Name); err != nil {
			apis.BadRequestErr(ctx, err)
			return
		}
		if err := validate.IsValidateEmail(user.Email); err != nil {
			apis.BadRequestErr(ctx, err)
			return
		}

		if err := db.Create(user).Error; err != nil {
			apis.InternalServerErr(ctx, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, user)
	})
	if err != nil {
		logrus.Errorf("%s", err)
		apis.InternalServerErr(ctx, err)
		return
	}
}

// func ListUsers(ctx *gin.Context) {
// 	us, err := db.ListUsers(ListUserOption{})
// 	if err != nil {
// 		apis.InternalServerErr(ctx, err)
// 		return
// 	}
// 	logrus.Debugf("list users")

// 	ctx.JSON(http.StatusOK, us)
// }

// func GetUserProfile(ctx *gin.Context) {
// 	uname := ctx.Param("name")
// 	u, err := db.GetUser(uname)
// 	if err != nil {
// 		apis.InternalServerErr(ctx, err)
// 		return
// 	}
// 	logrus.Debugf("%s: %+v", ctx.Request.URL.String(), u)
// 	ctx.JSON(http.StatusOK, u)
// }
