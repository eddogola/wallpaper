package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/eddogola/wallpapr"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	client  = &wallpapr.Client{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wallpapr",
	Short: "wallpapr sets a slideshow on the background of your desktop",
	Long: `Wallpapr gets photos from Unsplash, based on your query - 
a certain user's photos, a topic's photos, a search query to unsplash, 
or the first page's photos.`,
	PostRun: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		exitOnError(err)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wallpapr.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wallpapr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wallpapr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	auth, err := wallpapr.ReadAuthKeys()
	if err != nil {
		client.Logger.Println("error trying to read authentication keys from environment variables")
		os.Exit(1)
	}

	// set up http client
	// not so great - disable client side certificate verification

	cl := &http.Client{
		Transport: &http.Transport{
			// TLSClientConfig: &tls.Config{},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client = wallpapr.NewClient(auth)
	client.HTTPClient = cl
}
