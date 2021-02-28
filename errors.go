package wallpaper

import (
	"errors"
)

var (
	errKeysNotInEnv = errors.New("authentication keys not found in os env")
)
