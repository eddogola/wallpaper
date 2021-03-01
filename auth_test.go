package wallpaper

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestUnsplashAuth(t *testing.T) {
	uAccessKey := "-knJmOPLdhquiCDza-henc9esz9-GHNJLOPS567qaJ"
	uSecretKey := "suhfjwehfKhJDAdjkd-afjuhfjqe99123nddfdfkafk"

	t.Run("env file data provided", func(t *testing.T) {
		want := &UnsplashAuth{
			AccessKey: uAccessKey,
			SecretKey: uSecretKey,
		}
		envData := fmt.Sprintf(`export UNSPLASH_ACCESS_KEY=%v
					export UNSPLASH_SECRET_KEY=%v`, uAccessKey, uSecretKey)
		got, err := ReadAuthKeys([]byte(envData))

		checkError(err, t)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("env file not provided, environment variables provided", func(t *testing.T) {
		os.Setenv("UNSPLASH_ACCESS_KEY", uAccessKey)
		os.Setenv("UNSPLASH_SECRET_KEY", uSecretKey)
		want := &UnsplashAuth{
			AccessKey: uAccessKey,
			SecretKey: uSecretKey,
		}

		got, err := ReadAuthKeys([]byte(``))

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