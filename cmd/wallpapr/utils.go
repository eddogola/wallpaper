package main

import (
	"os"
)

func createDir(filePath string) error {
	return os.MkdirAll(filePath, 0770)
}

func fileExists(filePath string) bool {
	if info, err := os.Stat(filePath); os.IsNotExist(err) || info.IsDir() {
		return false
	}

	return true
}
