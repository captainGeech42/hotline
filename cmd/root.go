package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hotline",
	Short: "HTTP/DNS request logging",
	Long: `Web app to log HTTP and DNS requests for research or testing purposes.
Source available at https://github.com/captainGeech42/hotline`,
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "~/.hotline.yml", "Path to config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
