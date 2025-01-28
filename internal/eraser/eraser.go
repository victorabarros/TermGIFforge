package eraser

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/victorabarros/termgifforge/pkg/models"
)

// Clean erase GIFs older than 24 hours
func Clean(statuses models.StatusDetails, lastAccess map[string]time.Time) {
	for id := range statuses.Status {
		if id == "waiting" || id == "error" || id == "invalid" {
			continue
		}

		access, ok := lastAccess[id]
		if !ok {
			log.Printf("entry '%s' is missing lastAccess \n", id)
			lastAccess[id] = time.Now().Add(-12 * time.Hour)
			continue
		}

		// TTL 24 hours
		yesterday := time.Now().Add(-24 * time.Hour)
		if access.Before(yesterday) {
			path := fmt.Sprintf("output/%s.gif", id)
			if err := os.Remove(path); err != nil {
				log.Printf("fail to remove '%s': %+2v\n", path, err)
			}
			// TODO use mutex here ?
			delete(lastAccess, id)
		}
	}
}
