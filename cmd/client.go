package cmd

import (
	"github.com/captainGeech42/hotline/internal/client"
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/spf13/cobra"
)

var (
	CallbackName   string
	ShowHistorical bool
)

func init() {
	ClientCommand.Flags().StringVarP(&CallbackName, "name", "n", "", "Existing callback name to use (leave blank to generate a new one)")
	ClientCommand.Flags().BoolVarP(&ShowHistorical, "show-historical", "a", false, "Show all previous callbacks")
	rootCmd.AddCommand(ClientCommand)
}

var ClientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run the hotline client",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(configPath)
		if cfg == nil {
			return
		}

		client.StartClient(CallbackName, ShowHistorical, cfg)
	},
}
