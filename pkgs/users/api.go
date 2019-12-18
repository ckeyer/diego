package users

import (
	"net/http"

	"github.com/ckeyer/diego/pkgs/apis/validate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CheckNamespace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exi, err := stogr.ExistsNamespace(ctx.Param("name"))
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, map[string]bool{"message": exi})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := &User{}
		if err := decodeBody(ctx, user); err != nil {
			BadRequestErr(ctx, err)
			return
		}
		if err := validate.IsDNS1035Label(user.Name); err != nil {
			BadRequestErr(ctx, err)
			return
		}
		if err := validate.IsValidateEmail(user.Email); err != nil {
			BadRequestErr(ctx, err)
			return
		}

		if err := stogr.CreateUser(user); err != nil {
			InternalServerErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func CreateOrg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		org := &Org{}
		if err := decodeBody(ctx, org); err != nil {
			BadRequestErr(ctx, err)
			return
		}
		if err := validate.IsDNS1035Label(org.Name); err != nil {
			BadRequestErr(ctx, err)
			return
		}

		if err := stogr.CreateOrg(org); err != nil {
			InternalServerErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, org)
	}
}

func ListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		us, err := stogr.ListUsers(ListUserOption{})
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		logrus.Debugf("list users")

		ctx.JSON(http.StatusOK, us)
	}
}

func GetUserProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uname := ctx.Param("name")
		u, err := stogr.GetUser(uname)
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		logrus.Debugf("%s: %+v", ctx.Request.URL.String(), u)
		ctx.JSON(http.StatusOK, u)
	}
}

func ListOrgs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		os, err := stogr.ListOrgs(ListOrgOption{})
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, os)
	}
}

func GetOrgProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
