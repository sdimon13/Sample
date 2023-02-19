package cmd

import (
	"git.sample.ru/sample/internal/application"
	"git.sample.ru/sample/internal/exithandler"
	"git.sample.ru/sample/internal/logger"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API mode for application",
	Long:  `API mode for application startup`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := application.Get(application.TypeApi)
		if err != nil {
			logger.Error.Fatal(err.Error())
		}

		go func() {
			logger.Info.Printf("starting GRPC server at %s", app.Cfg.GrpcPort)
			if err := app.Api.Run(); err != nil {
				logger.Error.Fatal(err.Error())
			}
		}()

		exithandler.Init(func() {
			if err := app.Stop(); err != nil {
				logger.Error.Println(err)
			}

			if err := app.Api.Close(); err != nil {
				logger.Error.Println(err.Error())
			}
		})

	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
