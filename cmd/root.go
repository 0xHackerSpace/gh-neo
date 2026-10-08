package cmd

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "neo",
	Short: "gh-neo is a GitHub CLI extension",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.DefaultRESTClient()
		if err != nil {
			return err
		}

		response := struct{ Login string }{}
		if err := client.Get("user", &response); err != nil {
			return err
		}

		fmt.Printf("running as %s\n", response.Login)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
