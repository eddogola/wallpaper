package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var downloadCmd = cobra.Command{
	Use:     "download",
	Short:   "downloads a random photo and saves it to the provided directory",
	Example: "wallpapr download -o picture.jpeg -l ~/Downloads/wallpapers",
	Run: func(cmd *cobra.Command, args []string) {
		err := download()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	//fileName string
	location string
)

func init() {
	rootCmd.AddCommand(&downloadCmd)

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
