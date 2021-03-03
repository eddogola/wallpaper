package main

import (
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
)

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

var downloadCmd = cobra.Command{
	Use:     "download",
	Short:   "downloads a random photo and saves it to the provided directory",
	Example: "wallpapr download -o picture.jpeg -l ~/Downloads/wallpapers",
	Run: func(cmd *cobra.Command, args []string) {
		//create `location` if it doesn't exist
		if !fileExists(location) {
			createDir(location)
		}

		err := download()
		exitOnError(err)
	},
}

var (
	//fileName string
	location string
)

func init() {
	rootCmd.AddCommand(&downloadCmd)

	homeDir, err := getHomeDir()
	exitOnError(err)
	defaultDlLocation := homeDir + string(os.PathSeparator) + strings.Join([]string{"Downloads", "wallpapers"}, string(os.PathSeparator))

	/*
		implement the feature of providing custom file names later
		I'll have to change my downloader(clipr) drastically
	*/
	//downloadCmd.Flags().StringVar(&fileName, "o", "", "specify the filename of the downloaded file.")
	downloadCmd.Flags().StringVarP(&location, "location", "l", defaultDlLocation, "specify the directory where to save downloaded files.")
}

func download() error {
	pic, err := client.GetRandomPhoto()
	if err != nil {
		return err
	}
	pic.Download(location)

	return nil
}
