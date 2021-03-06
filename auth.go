package wallpapr

import (
	"os"
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
func ReadAuthKeys() (*UnsplashAuth, error) {
	// read auth keys from environment
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
