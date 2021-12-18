package cmd

import (
	"github.com/spf13/cobra"
)

var (
	CallbackName   string
	ShowHistorical bool
)

func init() {
	ClientCommand.Flags().StringVarP(&CallbackName, "name", "n", "", "Callback name to use (leave blank to generate a randome one)")
	ClientCommand.Flags().BoolVarP(&ShowHistorical, "show-historical", "a", false, "Show all previous callbacks")
	rootCmd.AddCommand(ClientCommand)
}

var ClientCommand = &cobra.Command{
	Use:   "client",
	Short: "Run the hotline client",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
