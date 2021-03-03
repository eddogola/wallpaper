package wallpapr

import (
	"fmt"
	"os/exec"
)

/*
set desktop background using gsettings

$ gsettings set org.gnome.desktop.background picture-uri 'file:///home/JohnDoe/Pictures/wallpaper.jpg'

*/

// SetWallpaper uses gsettings to set background pic
func SetWallpaper(filePath string) error {
	file := fmt.Sprintf("file://%v", filePath)
	gsettingsCmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", file)

	if err := gsettingsCmd.Run(); err != nil {
		return fmt.Errorf("please check your installation contains gsettings, err: %w", err)
	}

	return nil
}
