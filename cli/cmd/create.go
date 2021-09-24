package cmd

import (
	"fmt"
	"os"

	"github.com/mrityunjaygr8/shorty/app"
	"github.com/mrityunjaygr8/shorty/utils"

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
		config, err := utils.GetConfig()
		if err != nil {
			return err
		}
		app := app.Setup(config)
		token, err := app.Create(args[0])
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, token)

		return nil
	},
}
