package wallpaper

import (
	"net/http"
)

// Client defines methods to communicate with Unsplash's API
type Client struct {
	Auth       *UnsplashAuth
	HTTPClient *http.Client
	Config     *Config
}

// Config defines settings used in accessing the API
type Config struct {
	ItemsPerPage  int  // 10 - 30
	ContentFilter bool // to give flexibility in filtering content further; default is low
}

// NewClient creates a new Client with provided authentication keys
func NewClient(auth *UnsplashAuth) *Client {
	return &Client{Auth: auth}
}
