package cmd

import (
	"github.com/gleo-systems/retromeet/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start application server",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewInstance()
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
