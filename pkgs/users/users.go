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

const (
	TUsers = "users"
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
	return TUsers
}

type ListUsersOption struct {
	apis.ListOption
}

// CreateUser 新建用户
// @Tags 用户
// @Summary 新建
// @Description 新建用户
// @Produce json
// @Param user body users.User true "用户信息"
// @Success 201
// @Router /api/users [POST]
func CreateUser(ctx *gin.Context) {
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

	apis.GinInvoke(ctx, func(db *gorm.DB) {
		if err := db.Create(user).Error; err != nil {
			apis.InternalServerErr(ctx, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
}

type UserList struct {
	apis.ListOption
	Items []*User `json:"items"`
}

// ListUsers 用户列表
// @Tags 用户
// @Summary 列表
// @Description 用户列表
// @Produce json
// @Param offset query int false "起始位置"
// @Param limit query int false "数量"
// @Success 200 {object} UserList
// @Router /api/users [GET]
func ListUsers(ctx *gin.Context) {
	opt := &ListUsersOption{}
	if err := apis.Query(ctx, opt); err != nil {
		apis.BadRequestErr(ctx, err)
		return
	}
	apis.GinInvoke(ctx, func(db *gorm.DB) {
		ret := &UserList{
			Items: []*User{},
		}
		query := db.Table(TUsers)

		if err := query.Count(&ret.Count).Error; err != nil {
			apis.InternalServerErr(ctx, err)
			return
		}

		err := query.Offset(opt.Offset).
			Limit(opt.Limit).
			Find(&ret.Items).
			Error
		if err != nil {
			apis.InternalServerErr(ctx, err)
			return
		}
		logrus.Debugf("list users")

		ctx.JSON(http.StatusOK, ret)
	})
}

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
