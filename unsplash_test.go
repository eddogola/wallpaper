package wallpaper

import (
	"fmt"
	"testing"
)

func TestUnsplashAuth(t *testing.T) {
	t.Run("env file data provided", func(t *testing.T) {
		uAccessKey := "-knJmOPLdhquiCDza-henc9esz9-GHNJLOPS567qaJ"
		uSecretKey := "suhfjwehfKhJDAdjkd-afjuhfjqe99123nddfdfkafk"
		want := &UnsplashAuth{
			AccessKey: uAccessKey,
			SecretKey: uSecretKey,
		}
		envData := fmt.Sprintf(`UNSPLASH_ACCESS_KEY=%v
					UNSPLASH_SECRET_KEY=%v`, uAccessKey, uSecretKey)
		got, err := readAuthKeys([]byte(envData))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if want.AccessKey != got.AccessKey || want.SecretKey != got.SecretKey {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
