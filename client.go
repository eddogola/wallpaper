package wallpaper

import (
	"context"
	"fmt"
	"io/ioutil"
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

func (c *Client) getHTTP(ctx context.Context, link string) (*http.Response, error) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request")
	}

	authorizationHeaderValue := fmt.Sprintf("Client-ID %s", c.Auth.AccessKey)
	req.Header.Set("Authorization", authorizationHeaderValue)
	queryVars := req.URL.Query()
	queryVars.Add("per_page", fmt.Sprint(c.Config.ItemsPerPage))
	req.URL.RawQuery = queryVars.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request %v", req)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errUnexpectedStatusCode(resp.StatusCode)
	}
	return resp, nil
}

func (c *Client) getHTTPBodyBytes(ctx context.Context, link string) ([]byte, error) {
	resp, err := c.getHTTP(ctx, link)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
