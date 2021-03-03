package wallpapr

import (
	"github.com/eddogola/clipr"
)

// Download saves the picture to the provided location
func (pic *Photo) Download(location string) (string, error) {
	downloadURL := pic.Links.Download
	rs := &clipr.Resource{URL: downloadURL}
	err := rs.SetUp()
	if err != nil {
		return "", err
	}
	return rs.Download(location, pic.ID)
}
