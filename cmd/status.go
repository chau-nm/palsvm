package cmd

import (
	"fmt"

	"github.com/chau-nm/palsvm/internal/server"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if server.IsRunning() {
			fmt.Println("Server is running")
		} else {
			fmt.Println("Server is not running")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
