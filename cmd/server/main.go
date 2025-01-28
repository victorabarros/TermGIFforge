package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/termgifforge/internal/eraser"
	"github.com/victorabarros/termgifforge/internal/files"
	"github.com/victorabarros/termgifforge/internal/gif"
	"github.com/victorabarros/termgifforge/internal/id"
	"github.com/victorabarros/termgifforge/pkg/models"
)

var (
	GIFStatuses = struct {
		Fail       models.GIFStatus
		Processing models.GIFStatus
		Ready      models.GIFStatus
	}{
		Fail:       models.GIFStatus("Fail"),
		Processing: models.GIFStatus("Processing"),
		Ready:      models.GIFStatus("Ready"),
	}
	port = "80"

	setCmds = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	statuses   = map[string]models.GIFStatus{}
	lastAccess = map[string]time.Time{}
)

func init() {
	if err := files.CreateOutputDirectory(); err != nil {
		os.Exit(1)
	}

	gifs, err := files.ListGIFs()
	if err != nil {
		os.Exit(1)
	}
	for _, gif := range gifs {
		name := gif.Name()
		// remove .gif from name
		id := name[:len(name)-4]
		// TODO mutex
		statuses[id] = GIFStatuses.Ready
	}

	// TODO mutex
	if status := statuses["error"]; status != GIFStatuses.Ready {
		errorGIF()
	}
	// TODO mutex
	if status := statuses["invalid"]; status != GIFStatuses.Ready {
		invalidGIF()
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/victorabarros/TermGIFforge?tab=readme-ov-file#termgifforge-")
	})

	rpcGroup := r.Group("/api/v1")
	rpcGroup.GET("/gif", GetTerminalGIF)
	rpcGroup.GET("/mock", func(c *gin.Context) {
		c.File("output/error.gif")
	})

	// go cleaner() // TODO validade before enable it in production
	log.Println("Starting app in port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Printf("%+2v/n", err)
	}
}

func cleaner() {
	sleepLapse := 1 * time.Hour
	if os.Getenv("ENVIRONMENT") == "local" {
		sleepLapse = 1 * time.Second
	}

	for {
		time.Sleep(sleepLapse)
		// TODO mutex
		eraser.Clean(statuses, lastAccess)
	}
}

func GetTerminalGIF(c *gin.Context) {
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Printf("Error trying to serialize object: %+2v\n", err)
		c.File("output/invalid.gif")
		return
	}

	inputHash := id.NewUUUIDAsString(cmdsInputStr)
	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)
	// TODO mutex
	if status, ok := statuses[inputHash]; ok {
		if status == GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if status == GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if status == GIFStatuses.Ready {
			// TODO mutex
			lastAccess[inputHash] = time.Now()
			c.File(outGifPath)
			return
		}
	}

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	// TODO mutex
	statuses[id] = GIFStatuses.Processing

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		// TODO mutex
		statuses[id] = GIFStatuses.Fail
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		// TODO mutex
		statuses[id] = GIFStatuses.Fail
		return err
	}

	// TODO mutex
	statuses[id] = GIFStatuses.Ready
	// TODO log GIF done
	return nil
}

func errorGIF() error {
	cmdInput := []string{
		"Type \"Sorry, it was not possible create your GIF. =/\"",
		"Sleep 6s",
	}

	inputHash := "error"

	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	return nil
}

func invalidGIF() error {
	cmdInput := []string{
		"Type \"Invalid request...\"",
		"Sleep 6s",
	}

	inputHash := "invalid"

	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	return nil
}
