package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/eddogola/wallpapr"
	"github.com/spf13/cobra"
)

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

var downloadSetCmd = cobra.Command{
	Use:     "set-random",
	Short:   "downloads a random photo, saves it in the provided directory & sets it as the desktop background",
	Example: "wallpapr set-random -o picture.jpeg -l ~/Downloads/wallpapers",
	Run: func(cmd *cobra.Command, args []string) {
		//create `location` if it doesn't exist
		if !fileExists(location) {
			createDir(location)
		}

		err := setRandom()
		exitOnError(err)
	},
}

var (
	//fileName string
	location string
)

func init() {
	rootCmd.AddCommand(&downloadSetCmd)

	homeDir, err := getHomeDir()
	exitOnError(err)
	defaultDlLocation := homeDir + string(os.PathSeparator) + strings.Join([]string{"Downloads", "wallpapers"}, string(os.PathSeparator))

	/*
		implement the feature of providing custom file names later
		I'll have to change my downloader(clipr) drastically
	*/
	//downloadCmd.Flags().StringVar(&fileName, "o", "", "specify the filename of the downloaded file.")
	downloadSetCmd.Flags().StringVarP(&location, "location", "l", defaultDlLocation, "specify the directory where to save downloaded files.")
}

func setRandom() error {
	pic, err := client.GetRandomPhoto()
	if err != nil {
		return err
	}
	fn, err := pic.Download(location)
	if err != nil {
		return err
	}
	err = wallpapr.SetWallpaper(fn)
	if err != nil {
		return err
	}
	fmt.Printf("%v set as wallpaper\n", fn)

	return nil
}
