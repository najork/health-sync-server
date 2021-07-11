package cmd

import (
	"github.com/najork/apollo-server/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the apollo server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.New().Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
