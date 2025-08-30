package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/victorabarros/termgifforge/pkg/models"
)

var (
	sleepLapse        = 10 * time.Minute
	defaultLastAccess = time.Now().Add(-12 * time.Hour)
	ttl               = -24 * time.Hour
)

// CreateOutputDirectory creates ./output directory if it doesn't exist
func CreateOutputDirectory() error {
	dirName := "output"

	// Check if the directory exists
	_, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the directory if it doesn't exist
			perm := os.FileMode(0755)
			if err := os.Mkdir(dirName, perm); err != nil {
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

// ListGIFs returns slice of GIFs as os.DirEntry from ./output directory
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

// Cleaner is a worker that every hour removes GIFs older than TTL
func Cleaner(details *models.GIFDetails) {
	for {
		time.Sleep(sleepLapse)
		log.Println("Init cleaner")
		for id := range details.GIF {
			EraseGIF(id, details)
		}

		// TODO check if volume of GIF is higher than 90% of the DISK. If so, erase quarter of the oldest
	}
}

// EraseGIF clean GIF by id
func EraseGIF(id string, details *models.GIFDetails) {
	if id == "waiting" || id == "error" || id == "invalid" {
		return
	}

	d, ok := details.Get(id)
	if !ok {
		// 12 hour default last access
		details.SetLastAccess(id, defaultLastAccess)
		return
	}

	// Only remove if last access is older than TTL
	if d.LastAccess.Before(time.Now().Add(ttl)) {
		path := fmt.Sprintf("output/%s.gif", id)
		log.Printf("removing GIF %s \n", path)
		if err := os.Remove(path); err != nil {
			log.Printf("fail to remove '%s': %+2v\n", path, err)
			return
		}

		details.Del(id)
	}
}
