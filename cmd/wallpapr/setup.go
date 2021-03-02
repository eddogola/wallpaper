package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/eddogola/wallpapr"
)

/*
sets up the Client and Downloader to be used in the requests
to the Unsplash API and picture downloading
*/

var defaultDlLocation = "~/Downloads/wallpapers"
var client = &wallpapr.Client{}

func init() {
	data, err := readEnvFile(envFile)

	auth, err := wallpapr.ReadAuthKeys(data)
	exitOnError(err)

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

func readEnvFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
