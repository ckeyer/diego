package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	rootCmd.AddCommand(docsCmd())
}

// @title Diego API 文档
// @version v1.0
// @description 版本发布程序

// @contact.name @ckeyer
// @contact.url https://www.ckeyer.com/
// @contact.email me@ckeyer.com

// @host diego.ckeyer.com
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Token
func docsCmd() *cobra.Command {
	var addr string
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "文档api",
		Run: func(c *cobra.Command, args []string) {
			r := gin.New()
			r.GET("/api/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
			r.Run(addr)
		},
	}
	cmd.Flags().StringVar(&addr, "addr", ":8081", "api address.")
	return cmd
}
