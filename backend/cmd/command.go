package cmd

import (
	"fmt"
	"git.sample.ru/sample/internal/application"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/pkg/golibs/command"
	"github.com/spf13/cobra"
)

// commandCmd represents the command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "run command",
	Long:  `Run one-time command`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := map[string]string{
			"migrate":        "Run migration scripts",
			"migrate-create": "Creates new migration script",
			"migrate-status": "Returns migration status",
			"sync":           "Sync data from old source",
		}

		if len(args) == 0 || len(args) > 2 {
			fmt.Println("available commands")
			fmt.Println("+---------------------+-----------------------------------------------------------------+")
			for s, d := range commands {
				if len(d) > 64 {
					d = d[:61] + "..."
				}
				fmt.Printf("| %-20s| %-64s|\n", s, d)
			}
			fmt.Println("+---------------------+-----------------------------------------------------------------+")
			return
		}
		c := args[0]

		if _, ok := commands[c]; !ok {
			logger.Error.Fatal("command not available")
		}

		app, err := application.Get(application.TypeCommand)
		if err != nil {
			logger.Error.Fatal(err.Error())
		}

		switch c {
		case "migrate":
			logger.Info.Println("running migrations")
			c, err := command.NewMigrate(app.Cfg.DSN, app.Cfg.DbSchema, app.Cfg.MigrationPath)
			if err != nil {
				logger.Error.Fatal(err)
			}
			err = c.Up()
			if err != nil {
				logger.Error.Fatal(err)
			}
		case "migrate-status":
			c, err := command.NewMigrate(app.Cfg.DSN, app.Cfg.DbSchema, app.Cfg.MigrationPath)
			if err != nil {
				logger.Error.Fatal(err)
			}

			err = c.Status()
			if err != nil {
				logger.Error.Fatal(err)
			}
		case "migrate-create":
			c, err := command.NewMigrate(app.Cfg.DSN, app.Cfg.DbSchema, app.Cfg.MigrationPath)
			if err != nil {
				logger.Error.Fatal(err)
			}

			var name string
			if len(args) == 2 {
				name = args[1]
			} else {
				fmt.Println("Enter migration name:")
				_, err := fmt.Scanln(&name)
				if err != nil {
					logger.Error.Fatal(err)
				}
			}
			err = c.Create(name)
			if err != nil {
				logger.Error.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(commandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
