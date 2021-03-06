package main

import (
	"io/ioutil"
	"os"
)

func readEnvFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createDir(filePath string) error {
	return os.MkdirAll(filePath, 0770)
}

func fileExists(filePath string) bool {
	if info, err := os.Stat(filePath); os.IsNotExist(err) || info.IsDir() {
		return false
	}

	return true
}
