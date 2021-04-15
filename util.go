package wallpapr

import (
	"fmt"
	"strings"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

func filenameFromPhoto(pic *client.Photo) string {
	firstName := pic.User.FirstName
	lastName := pic.User.LastName
	ID := pic.ID
	name := strings.Join([]string{firstName, lastName, ID, "unsplash"}, "-")
	ext := "jpg"

	return fmt.Sprintf("%s.%s", name, ext)
}

func formPath(dst, filename string) string {
	return dst + filename
}
