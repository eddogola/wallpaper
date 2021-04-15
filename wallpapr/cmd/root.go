package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Wallpapr",
	Short: "Wallpapr sets wallpapers from unsplash on Debian devices",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}

// Execute runs rootCmdd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
