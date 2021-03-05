package main

import (
	"os"
	"strings"
	"time"

	"github.com/eddogola/wallpapr"
	"github.com/spf13/cobra"
)

var frequency time.Duration

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the desktop background",
	Long: `Batch downloads several photos requested by the user(by search query, topic or Unsplash user).
Starts a slideshow of photos in the photos' storage directory, with the photos changing after the specified
duration`,
	Example: "wallpapr set --search food --freq mins --time 15",
	// Args: cobra.exactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setWallpaper()
	},
}

var (
	all       bool
	userName  string
	topic     string
	search    string
	freq      string
	timeTaken float64
)

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().BoolVarP(&all, "all", "a", true, "Get photos from a list of all Unsplash photos")
	setCmd.Flags().StringVarP(&userName, "username", "u", "", "The username of an Unsplash user, from whom to get wallpapers.")
	setCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic from which to get photos from unsplash.")
	setCmd.Flags().StringVarP(&search, "search", "s", "", "The search query whose results will provide wallpapers.")
	setCmd.Flags().StringVarP(&freq, "freq", "f", "hours", "The frequency with which to change the wallpaper(mins, hours). Default is `hours`.")
	setCmd.Flags().Float64Var(&timeTaken, "time", 20, "The time after which the wallpaper should change, in the given frequency. For example, --freq mins --time 10: the wallpaper will change every 10 minutes.")
}

func setWallpaper() {
	var photos []wallpapr.Photo
	if all {
		pics, err := client.GetFirstPagePhotos()
		exitOnError(err)
		photos = append(photos, pics...)
	}
	if userName != "" {
		pics, err := client.GetPhotosByUser(userName)
		exitOnError(err)
		photos = append(photos, pics...)
	}
	if topic != "" {
		pics, err := client.GetPhotosByTopic(topic)
		exitOnError(err)
		photos = append(photos, pics...)
	}
	if search != "" {
		pics, err := client.GetPhotosBySearchQuery(search)
		exitOnError(err)
		photos = append(photos, pics...)
	}

	homeDir, err := getHomeDir()
	exitOnError(err)
	defaultDlLocation := homeDir + string(os.PathSeparator) + strings.Join([]string{"Downloads", "wallpapers"}, string(os.PathSeparator))

	switch freq {
	case "mins":
		frequency = time.Minute * time.Duration(timeTaken)
	case "hours":
		frequency = time.Hour * time.Duration(timeTaken)
	}

	// set up slideshow
	slideshow := wallpapr.NewSlideshow(photos, frequency)
	slideshow.Rotate(defaultDlLocation)
}
