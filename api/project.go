package api

import (
	"encoding/json"
	"net/http"

	"github.com/ckeyer/diego/pkgs/apis/validate"
	"github.com/ckeyer/diego/types"
	"github.com/gin-gonic/gin"
)

func ListProjects() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		items, err := stogr.ListProjects(ctx.Param("namespace"))
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, items)
	}
}

func GetProjectProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item, err := stogr.GetProject(ctx.Param("namespace"), ctx.Param("name"))
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, item)
	}
}

func CreateProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prj := &types.Project{}
		err := json.NewDecoder(ctx.Request.Body).Decode(prj)
		if err != nil {
			BadRequestErr(ctx, err)
			return
		}

		if errstrs := validate.IsDNS1035Label(prj.Name); len(errstrs) > 0 {
			BadRequestErr(ctx, errstrs)
			return
		}

		prj.Namespace = ctx.Param("namespace")
		if err := stogr.CreateProject(prj); err != nil {
			InternalServerErr(ctx, err)
			return
		}

		ctx.Writer.WriteHeader(http.StatusNoContent)
	}
}

func RemoveProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := stogr.RemoveProject(ctx.Param("namespace"), ctx.Param("name"))
		if err != nil {
			InternalServerErr(ctx, err)
			return
		}
		ctx.Writer.WriteHeader(http.StatusNoContent)
	}
}

func UploadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo(ctx)
	}
}
