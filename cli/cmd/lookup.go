package cmd

import (
	"fmt"
	"github.com/mrityunjaygr8/shorty/app"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(lookupCmd)
}

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "A command to lookup the long URL for the short token",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		app := app.Setup(app.Config{DB_NAME: "postgres", DB_USER: "root", DB_PASS: "secret", DB_HOST: "localhost", DB_PORT: 5432, DB_SSL: "disable"})
		url, found, err := app.Lookup(args[0])
		if err != nil {
			return err
		}

		if !found {
			fmt.Fprintf(os.Stdout, "Token: %s does not have an associated URL", args[0])
		}

		fmt.Fprintln(os.Stdout, url)

		return nil

	},
}
