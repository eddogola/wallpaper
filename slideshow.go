package wallpapr

import (
	"time"
)

// Slideshow defines methods that will rotate the slideshow photos
// using the provided time frequency
type Slideshow struct {
	Photos    []Photo
	Frequency time.Duration
}

// NewSlideshow constructs a new slideshow type with photos
// in the slideshow
func NewSlideshow(photos []Photo, freq time.Duration) *Slideshow {
	return &Slideshow{Photos: photos, Frequency: freq}
}

// Rotate changes photos after the given frequency
func (s *Slideshow) Rotate(dlLocation string) {
	go func() {
		for _, photo := range s.Photos {
			time.Sleep(s.Frequency)
			fp, err := photo.Download(dlLocation)
			if err != nil {
				continue
			}
			err = SetWallpaper(fp)
			if err != nil {
				continue
			}
		}
	}()
}
