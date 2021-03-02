package wallpapr

import (
	"fmt"

	"github.com/eddogola/clipr"
)

// Download saves the picture to the provided location
func (pic *Photo) Download(location string) error {
	downloadURL := pic.Links.Download
	rs := &clipr.Resource{URL: downloadURL}
	err := rs.SetUp()
	if err != nil {
		return err
	}

	fn, err := rs.Download(location)

	fmt.Println(fn)

	return err
}
