package cmd

import (
	"github.com/chau-nm/palsvm/internal/server"
	"github.com/spf13/cobra"
)

const (
	DefaultPort = "8080"
	DefaultHost = "0.0.0.0"
)

var (
	port string
	host string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the HTTP server",
	Long:  `Start the HTTP server on specified port`,
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := server.New(host, port)

		if err := srv.Start(); err != nil {
			return err
		}

		srv.Wait()

		return srv.Stop()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&port, "port", "p", DefaultPort, "Port to run server on")
	startCmd.Flags().StringVarP(&host, "host", "H", DefaultHost, "Host to bind server to")
}
