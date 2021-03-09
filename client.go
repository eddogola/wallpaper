package wallpapr

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client defines methods to communicate with Unsplash's API
type Client struct {
	Auth       *UnsplashAuth
	HTTPClient *http.Client
	Config     *Config
	Logger     Logger
}

// Config defines settings used in accessing the API
type Config struct {
	ItemsPerPage  int    // 10 - 30
	ContentFilter string // to give flexibility in filtering content further; default is low
}

// NewClient creates a new Client with provided authentication keys
func NewClient(auth *UnsplashAuth) *Client {
	config := &Config{ItemsPerPage: 30, ContentFilter: "low"}
	return &Client{Auth: auth, Config: config}
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

	// filter only pictures with a landscape orientation
	// where the query parameter is accepted
	queryVars.Add("orientation", "landscape")
	req.URL.RawQuery = queryVars.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request %v", err)
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

/*
Client methods to get photos from Unsplash
I will just work with the first page
since, at least for the moment,
I don't need all those photos.
*/

// GetFirstPagePhotos returns a list of all Photos from sorted by `latest`
func (c *Client) GetFirstPagePhotos() ([]Photo, error) {
	endPoint := "https://api.unsplash.com/photos"
	data, err := c.getHTTPBodyBytes(context.Background(), endPoint)
	if err != nil {
		return nil, err
	}

	var photos []Photo
	err = parseJSON(data, &photos)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

// GetPhotosByUser returns a list of Photos uploaded by a specific user
func (c *Client) GetPhotosByUser(username string) ([]Photo, error) {
	endPoint := fmt.Sprintf("https://api.unsplash.com/%s/photos", username)
	data, err := c.getHTTPBodyBytes(context.Background(), endPoint)
	if err != nil {
		return nil, err
	}

	var photos []Photo
	err = parseJSON(data, &photos)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

// GetPhotosBySearchQuery returns a list of Photo results for a search query
func (c *Client) GetPhotosBySearchQuery(searchQuery string) ([]Photo, error) {
	endPoint := "https://api.unsplash.com/search/photos"
	url, err := url.Parse(endPoint)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("content_filter", c.Config.ContentFilter)
	q.Add("query", searchQuery)
	url.RawQuery = q.Encode()

	data, err := c.getHTTPBodyBytes(context.Background(), url.String())
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = parseJSON(data, &result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}

// GetPhotosByTopic returns a list of photos in the provided topic
func (c *Client) GetPhotosByTopic(topic string) ([]Photo, error) {
	endPoint := fmt.Sprintf("https://api.unsplash.com/topics/%s/photos", topic)
	data, err := c.getHTTPBodyBytes(context.Background(), endPoint)
	if err != nil {
		return nil, err
	}

	var photos []Photo
	err = parseJSON(data, &photos)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

// GetRandomPhoto gets a random picture from Unsplash
func (c *Client) GetRandomPhoto() (Photo, error) {
	endPoint := "https://api.unsplash.com/photos/random"
	data, err := c.getHTTPBodyBytes(context.Background(), endPoint)
	if err != nil {
		return Photo{}, err
	}

	var pic Photo
	err = parseJSON(data, &pic)
	if err != nil {
		return Photo{}, err
	}
	return pic, err
}

// utility function to parse json data to desired object
func parseJSON(data []byte, desiredObject interface{}) error {
	if err := json.Unmarshal(data, desiredObject); err != nil {
		return fmt.Errorf("error parsing json data: %v", err)
	}

	return nil
}
