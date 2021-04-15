// Package wallpapr sets a random photo from unsplash.com
// as the desktop background for (at least right now) computers
// using the Debian distros.
package wallpapr

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cavaliercoder/grab"

	"github.com/eddogola/unsplash-go/unsplash"
	"github.com/eddogola/unsplash-go/unsplash/client"
)

// Wallpapr contains the unsplash API client wrapper
type Wallpapr struct {
	unsplash *unsplash.Unsplash
	dst      string
	logger   *log.Logger
}

// NewWallpapr initializes a Wallpapr with a basic stdout logger
// and a basic Unsplash object, unauthorized for private unsplash
// actions.
func NewWallpapr(dst, clientID string) *Wallpapr {
	return &Wallpapr{
		unsplash: unsplash.New(client.New(clientID, nil, client.NewConfig())),
		dst:      dst,
		logger:   log.New(os.Stdout, "", 1),
	}
}

// Get saves the `Photo` provided in `Wallpapr.dst`
func (w *Wallpapr) Get(pic *client.Photo) (string, error) {
	client := grab.NewClient()
	req, err := grab.NewRequest(w.dst, pic.Links.Download)
	// set filename
	req.Filename = formPath(w.dst, filenameFromPhoto(pic))
	if err != nil {
		return "", err
	}
	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("\033[1AProgress %d / %d bytes (%d%%)\033[K\n",
				resp.BytesComplete(),
				resp.Size,
				int(100*resp.Progress()))

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	return resp.Filename, nil
}
