package wallpapr

import (
	"testing"

	"github.com/eddogola/readenv"
)

func TestWallpapr(t *testing.T) {
	envData, err := readenv.File(".env")
	if err != nil {
		t.Fatal(err)
	}
	w := NewWallpapr("/home/ogola/", envData["UNSPLASH_ACCESS_KEY"])
	err = w.SetRandom()

	if err != nil {
		t.Fatal(err)
	}
}
