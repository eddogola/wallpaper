package main

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use: "set",
	Short: "wallpaper sets the desktop background",
	Example: "wallpaper set --search food --freq mins --time 15",
	// Args: cobra.exactArgs(1),
	// Run: func(cmd*cobra.Command, args []string) {
		// stuff
	// },
}

var (
	all bool
	userName string
	topic string
	search string
	freq string
	time float64
)

func init() {
	rootCmd.AddCommand(setCmd)
	
	setCmd.Flags().BoolVarP(&all, "all-photos", "all", true, "Get photos from a list of all Unsplash photos")
	setCmd.Flags().StringVarP(&userName, "username", "u", "", "The username of an Unsplash user, from whom to get wallpapers.")
	setCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic from which to get photos from unsplash.")
	setCmd.Flags().StringVarP(&search, "search", "s", "", "The search query whose results will provide wallpapers.")
	setCmd.Flags().StringVarP(&freq, "freq", "f", "hours", "The frequency with which to change the wallpaper(mins, hours, days). Default is `hours`.")
	setCmd.Flags().Float64VarP(&time, "time", "t", 20, "The time after which the wallpaper should change, in the given frequency. For example, --freq mins --time 10: the wallpaper will chnage every 10 minutes.")
}