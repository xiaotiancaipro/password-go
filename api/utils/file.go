package utils

import (
	"io"
	"os"
)

type FileUtil struct{}

func (f FileUtil) ReadFile(file string) []byte {
	configFile, err := os.Open(file)
	if err != nil {
		log.Error("Error opening config file: %v\n", err)
		os.Exit(1)
	}
	defer configFile.Close()
	fileData, err := io.ReadAll(configFile)
	if err != nil {
		log.Error("Error reading config file: %v\n", err)
		os.Exit(1)
	}
	return fileData
}
