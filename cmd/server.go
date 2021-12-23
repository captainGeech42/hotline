package cmd

import (
	"log"
	"os"

	"github.com/captainGeech42/hotline/internal/callback/dns"
	"github.com/captainGeech42/hotline/internal/callback/http"
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/web"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ServerCommand)
}

var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the Hotline server (set $HOTLINE_APP to configure which server to run)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(configPath)
		if cfg == nil {
			return
		}

		hotlineApp := os.Getenv("HOTLINE_APP")

		switch hotlineApp {
		case "dns":
			log.Println("starting dns callback server")

			dns.StartServer(cfg)
		case "http":
			log.Println("starting http callback server")

			http.StartServer(cfg)
		case "web":
			log.Println("starting web server")

			web.StartApp(cfg)
		default:
			log.Fatalln("failed to launch server: you must set $HOTLINE_APP to dns, http, or web")
		}
	},
}
