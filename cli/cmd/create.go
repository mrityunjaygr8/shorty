package cmd

import (
	"fmt"
	"os"

	"github.com/mrityunjaygr8/shorty/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A command to create a short token from a long URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		app := app.Setup(app.Config{DB_NAME: "postgres", DB_USER: "root", DB_PASS: "secret", DB_HOST: "localhost", DB_PORT: 5432, DB_SSL: "disable"})
		token, err := app.Create(args[0])
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, token)

		return nil
	},
}
