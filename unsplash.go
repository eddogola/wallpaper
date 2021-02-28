package wallpaper

import (
	"os"
	"regexp"
	"strings"
)

// UnsplashAuth contains confidential keys with which to access the unsplash API
type UnsplashAuth struct {
	AccessKey string
	SecretKey string
}

const (
	// UnsplashAccessKey defines the key of the access key environment variable
	UnsplashAccessKey = "UNSPLASH_ACCESS_KEY"
	// UnsplashSecretKey defines the key of the secret key environment variable
	UnsplashSecretKey = "UNSPLASH_SECRET_KEY"
)

func readAuthKeys(envData []byte) (*UnsplashAuth, error) {
	var unsplashAccessKey string
	var unsplashSecretKey string

	// read auth keys from environment if no .env file is provided
	if len(envData) == 0 {
		unsplashAccessKey = os.Getenv(UnsplashAccessKey)
		unsplashSecretKey = os.Getenv(UnsplashSecretKey)
		if unsplashAccessKey == "" || unsplashSecretKey == "" {
			return nil, errKeysNotInEnv
		}

		uAuth := UnsplashAuth{
			AccessKey: unsplashAccessKey,
			SecretKey: unsplashSecretKey,
		}
		return &uAuth, nil
	}

	// parse environment values from .env file to UnsplashAuth struct
	lines := strings.Split(string(envData), "\n")
	regex := regexp.MustCompile(`(?:\w+[^=])(?:=)([a-zA-Z0-9_-]+)`)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		value := regex.FindStringSubmatch(line)[1]
		if strings.Contains(line, "ACCESS") {
			unsplashAccessKey = value
		} else if strings.Contains(line, "SECRET") {
			unsplashSecretKey = value
		}
	}
	return &UnsplashAuth{AccessKey: unsplashAccessKey, SecretKey: unsplashSecretKey}, nil
}
