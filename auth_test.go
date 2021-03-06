package wallpapr

import (
	"os"
	"reflect"
	"testing"
)

func TestUnsplashAuth(t *testing.T) {
	uAccessKey := "-knJmOPLdhquiCDza-henc9esz9-GHNJLOPS567qaJ"
	uSecretKey := "suhfjwehfKhJDAdjkd-afjuhfjqe99123nddfdfkafk"

	t.Run("environment variables provided", func(t *testing.T) {
		os.Setenv("UNSPLASH_ACCESS_KEY", uAccessKey)
		os.Setenv("UNSPLASH_SECRET_KEY", uSecretKey)
		want := &UnsplashAuth{
			AccessKey: uAccessKey,
			SecretKey: uSecretKey,
		}

		got, err := ReadAuthKeys()

		checkError(err, t)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func checkError(err error, t *testing.T) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
