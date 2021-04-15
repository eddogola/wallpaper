package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eddogola/wallpapr"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const unsplashAccessKeyEnvKey = "UNSPLASH_ACCESS_KEY"

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defaultDownloadLocation := filepath.Join(home, "wallpapers") + "/"
	destinationFlagUsageString := "the directory where downloaded photos will be saved"
	setRandomCmd.Flags().StringVarP(&dst, "destination", "d", defaultDownloadLocation, destinationFlagUsageString)

	rootCmd.AddCommand(setRandomCmd)
}

var (
	dst string

	setRandomCmd = &cobra.Command{
		Use:   "set-random",
		Short: "gets a random wallpaper from unsplash and sets it as the desktop background",
		RunE: func(cmd *cobra.Command, args []string) error {
			// create directory `mkdidr -p`
			// if the destination(`dst`) filepath does not exist
			if !fileExists(dst) {
				err := createDir(dst)
				if err != nil {
					return err
				}
			}
			// read unsplash access key from os environment variables
			unsplashAccessKey := os.Getenv(unsplashAccessKeyEnvKey)
			if unsplashAccessKey == "" {
				return fmt.Errorf("access key variable '%v' not found in os environment variables", unsplashAccessKeyEnvKey)
			}
			// initialize a `Wallpapr` object and set random wallpaper
			w := wallpapr.NewWallpapr(dst, unsplashAccessKey)
			err := w.SetRandom()
			if err != nil {
				return err
			}
			return nil
		},
	}
)

// createDir creates all directories in the provided filepath
// works like the unix command `$ mkdir -p <path>`
func createDir(filePath string) error {
	return os.MkdirAll(filePath, 0770)
}

// checks if a filepath exists
// returns false if it doesn't exist
// true otherwise
func fileExists(filePath string) bool {
	if info, err := os.Stat(filePath); os.IsNotExist(err) || info.IsDir() {
		return false
	}

	return true
}
