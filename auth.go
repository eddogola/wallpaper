package wallpaper

import (
	"fmt"
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

// ReadAuthKeys extracts authorization credentials and stores
// them in an UnsplashAuth object
func ReadAuthKeys(envData []byte) (*UnsplashAuth, error) {
	uAuth, err := parseEnvFileData(envData)

	// read auth keys from environment if no .env file is provided
	if err != nil {
		unsplashAccessKey := os.Getenv(UnsplashAccessKey)
		unsplashSecretKey := os.Getenv(UnsplashSecretKey)
		if unsplashAccessKey == "" || unsplashSecretKey == "" {
			return nil, errKeysNotInEnvVars
		}

		uAuth := UnsplashAuth{
			AccessKey: unsplashAccessKey,
			SecretKey: unsplashSecretKey,
		}
		return &uAuth, nil
	}

	return uAuth, nil
}

// parse environment values from .env file to UnsplashAuth struct
func parseEnvFileData(data []byte) (*UnsplashAuth, error) {
	var unsplashAccessKey string
	var unsplashSecretKey string

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("lines in env data less than 2")
	}
	regex := regexp.MustCompile(`(?:\w+[^=])(?:=)([a-zA-Z0-9_-]+)`)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		value := regex.FindStringSubmatch(line)[1]
		if strings.Contains(line, "ACCESS") {
			unsplashAccessKey = value
		} else if strings.Contains(line, "SECRET") {
			unsplashSecretKey = value
		} else {
			return nil, errKeysNotInEnvFile
		}
	}
	return &UnsplashAuth{AccessKey: unsplashAccessKey, SecretKey: unsplashSecretKey}, nil
}
