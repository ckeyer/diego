package projects

import (
	"net/http"

	"github.com/ckeyer/diego/pkgs/apis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	// TProjects 项目表
	TProjects = "projects"
)

// Project 项目
type Project struct {
	gorm.Model

	Name      string `json:"name" gorm:"column:name;type:varchar(16)"`
	Namespace string `json:"namespace" gorm:"column:namespace;type:varchar(16)"`
	Desc      string `json:"desc" gorm:"column:desc;type:varchar(255)"`
}

// ListProjects 项目列表
func ListProjects(ctx *gin.Context) {
	apis.GinInvoke(ctx, func(db *gorm.DB) {
		items := []*Project{}
		err := db.Table(TProjects).
			Where("namespace = ?", ctx.Param("namespace")).
			Find(&items).
			Error
		if err != nil {
			apis.InternalServerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, items)
	})
}

// func GetProjectProfile() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		item, err := stogr.GetProject(ctx.Param("namespace"), ctx.Param("name"))
// 		if err != nil {
// 			apis.InternalServerErr(ctx, err)
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, item)
// 	}
// }

// func CreateProject() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		prj := &Project{}
// 		err := json.NewDecoder(ctx.Request.Body).Decode(prj)
// 		if err != nil {
// 			apis.BadRequestErr(ctx, err)
// 			return
// 		}

// 		if errstrs := validate.IsDNS1035Label(prj.Name); len(errstrs) > 0 {
// 			apis.BadRequestErr(ctx, errstrs)
// 			return
// 		}

// 		prj.Namespace = ctx.Param("namespace")
// 		if err := stogr.CreateProject(prj); err != nil {
// 			apis.InternalServerErr(ctx, err)
// 			return
// 		}

// 		ctx.Writer.WriteHeader(http.StatusNoContent)
// 	}
// }

// func RemoveProject() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		err := stogr.RemoveProject(ctx.Param("namespace"), ctx.Param("name"))
// 		if err != nil {
// 			apis.InternalServerErr(ctx, err)
// 			return
// 		}
// 		ctx.Writer.WriteHeader(http.StatusNoContent)
// 	}
// }

// func UploadFile() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		todo(ctx)
// 	}
// }
