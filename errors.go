package wallpaper

import (
	"errors"
)

var (
	errKeysNotInEnvVars = errors.New("authentication keys not found in os env")
	errKeysNotInEnvFile = errors.New("authentication keys not found in provided .env file")
)
