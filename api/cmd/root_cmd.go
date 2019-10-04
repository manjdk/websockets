package cmd

import (
	"log"

	"github.com/manjdk/websockets/api/server"
	"github.com/spf13/cobra"
)

const defaultPort = 8080

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Websocket server",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		server.Run(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().IntP("port", "p", defaultPort, "Port to run on")
}
