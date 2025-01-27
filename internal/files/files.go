package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateOutputDirectory() error {
	dirName := "output"

	// Check if the directory exists
	_, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			if err := os.Mkdir(dirName, 0755); err != nil {
				log.Printf("Failed to create directory %s: %v\n", dirName, err)
				return err
			}
			return nil
		}

		fmt.Printf("Fail to check if Directory '%s' exists: %v\n", dirName, err)
		return err
	}

	return nil
}

func ListGIFs() ([]os.DirEntry, error) {
	dirName := "output"
	var gifFiles []os.DirEntry
	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %v", dirName, err)
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".gif" {
			gifFiles = append(gifFiles, file)
		}
	}

	return gifFiles, nil
}
