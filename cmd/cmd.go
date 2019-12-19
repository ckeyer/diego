package cmd

import (
	"github.com/ckeyer/diego/api"
	"github.com/ckeyer/diego/hacks/webhook"
	"github.com/ckeyer/diego/pkgs/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = cobra.Command{
		Use:   "diego",
		Short: "diego 版本发布系统",
		Run:   runServe,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			if app.Debug {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.Debug("debug mode")
			}
		},
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&app.Debug, "debug", "D", false, "debug level")
	rootCmd.Flags().StringVarP(&app.Addr, "addr", "s", ":8080", "web server listenning address.")
	rootCmd.Flags().StringVar(&app.DBType, "db-type", "postgres", "postgres OR mysql.")
	rootCmd.Flags().StringVar(&app.DBURL, "db-url", "", "database connect URL")
	rootCmd.Flags().StringVar(&webhook.HookFile, "hook-config", "tools/webhook/webhook.yaml", "webhook config file.")
}

// runServe start http server.
func runServe(cmd *cobra.Command, args []string) {
	logrus.Infof("listenning at %s", app.Addr)
	if err := api.Serve(app.Addr); err != nil {
		logrus.Error(err)
	}
}

// Execute cmd main
func Execute() {
	rootCmd.AddCommand(printVersionCmd())
	rootCmd.Execute()
}

func printVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			// version.Print(nil)
		},
	}
}
