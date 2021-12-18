package cmd

import (
	"log"
	"os"

	"github.com/captainGeech42/hotline/internal/app"
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
		case "http":
			log.Println("starting http callback server")
		case "app":
			log.Println("starting app server")

			app.StartApp(config)
		default:
			log.Fatalln("You must set $HOTLINE_APP to dns, http, or app")
		}
	},
}
