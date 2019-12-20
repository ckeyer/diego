package users

import (
	"net/http"
	"time"

	"github.com/ckeyer/diego/pkgs/apis"
	"github.com/ckeyer/diego/pkgs/apis/validate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type User struct {
	gorm.Model

	Name  string `json:"name" gorm:"column:name;type:varchar(16)"`
	Email string `json:"email" gorm:"column:email;type:varchar(32)"`
	Desc  string `json:"desc" gorm:"column:desc;type:varchar(255)"`

	Password *string `json:"-" gorm:"column:password;type:varchar(64)"`

	FirstLoginAt  *time.Time `json:"first_login_at" gorm:"column:first_login_at"`
	LastLoginAt   *time.Time `json:"last_login_at" gorm:"column:last_login_at"`
	LastLoginIP   *string    `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(16)"`
	LastLoginType *string    `json:"last_login_type" gorm:"column:last_login_type;type:varchar(8)"`
}

func (User) TableName() string {
	return "users"
}

type ListUserOption struct {
	apis.ListOption
}

func CreateUser(ctx *gin.Context) {
	apis.GinInvoke(ctx, func(db *gorm.DB) {
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