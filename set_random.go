package wallpapr

import (
	"github.com/eddogola/unsplash-go/unsplash/client"
)

// SetRandom gets a random photo from `Wallpapr.unsplash.Photos.Random`
// downloads and saves the photo to `Wallpapr.dst`
// and sets the photo aas the desktop background.
func (w *Wallpapr) SetRandom() error {
	// get random photo using unsplash api client
	resp, err := w.unsplash.Photos.Random(nil)
	if err != nil {
		return err
	}
	pic := resp.(*client.Photo)

	// download photo using grab
	fn, err := w.Get(pic)
	if err != nil {
		return err
	}

	err = SetWallpaper(fn)
	if err != nil {
		return err
	}
	w.logger.Printf("%v set as wallpaper\n", fn)

	return nil
}
