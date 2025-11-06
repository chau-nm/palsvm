package cmd

import (
	"github.com/chau-nm/palsvm/internal/server"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.StopRunningServer()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
