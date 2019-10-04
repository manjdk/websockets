package cmd

import (
	"log"

	"github.com/manjdk/websockets/client/connection"
	"github.com/spf13/cobra"
)

const defaultPort = 8080

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Websocket client",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		connection.Dial(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().IntP("port", "p", defaultPort, "Port to run client on")
}
