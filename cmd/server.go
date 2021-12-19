package cmd

import (
	"log"
	"os"

	"github.com/captainGeech42/hotline/internal/app"
	"github.com/captainGeech42/hotline/internal/callback/dns"
	"github.com/captainGeech42/hotline/internal/callback/http"
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ServerCommand)
}

var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the hotline server (set $HOTLINE_APP to configure which server to run)",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig(configPath)

		hotlineApp := os.Getenv("HOTLINE_APP")

		switch hotlineApp {
		case "dns":
			log.Println("starting dns callback server")

			dns.StartServer(config)
		case "http":
			log.Println("starting http callback server")

			http.StartServer(config)
		case "app":
			log.Println("starting app server")

			app.StartApp(config)
		default:
			log.Fatalln("failed to launch server: you must set $HOTLINE_APP to dns, http, or app")
		}
	},
}
