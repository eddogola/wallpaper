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
