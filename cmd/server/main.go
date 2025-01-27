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
	"github.com/victorabarros/termgifforge/pkg/model"
)

var (
	GIFStatuses = struct {
		Fail       model.GIFStatus
		Processing model.GIFStatus
		Ready      model.GIFStatus
	}{
		Fail:       model.GIFStatus("Fail"),
		Processing: model.GIFStatus("Processing"),
		Ready:      model.GIFStatus("Ready"),
	}
	port = "80"

	setCmds = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	statuses   = map[string]model.GIFStatus{}
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
		id := name[:len(name)-4]
		statuses[id] = GIFStatuses.Ready
	}

	if status := statuses["waiting"]; status != GIFStatuses.Ready {
		waitingGIF()
	}
	if status := statuses["error"]; status != GIFStatuses.Ready {
		errorGIF()
	}
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
		// c.File("output/waiting.gif")
	})

	// go cleaner() // TODO validade before enable it in production
	log.Println("Starting app in port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Printf("%+2v/n", err)
	}
}

func cleaner() {
	sleepLapse := 24 * time.Hour
	if os.Getenv("ENVIRONMENT") == "local" {
		sleepLapse = 1 * time.Second
	}

	for {
		time.Sleep(sleepLapse)
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
	if status, ok := statuses[inputHash]; ok {
		if status == GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if status == GIFStatuses.Processing {
			c.File("output/waiting.gif")
			return
		}
		if status == GIFStatuses.Ready {
			lastAccess[inputHash] = time.Now()
			c.File(outGifPath)
			return
		}
	}

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	c.File("output/waiting.gif")
}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	statuses[id] = GIFStatuses.Processing

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		statuses[id] = GIFStatuses.Fail
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		statuses[id] = GIFStatuses.Fail
		return err
	}

	statuses[id] = GIFStatuses.Ready

	return nil
}

func waitingGIF() error {
	msg := "Wait..."
	cmdInput := []string{
		"Set FontSize 15",
		"Type \"PROCESSING_YOUR_GIF=true\"", "Enter", "Sleep 250ms",
		"Type \"while $PROCESSING_YOUR_GIF; do\"", "Enter", "Sleep 250ms",
		fmt.Sprintf("Type \"   echo '%s'\"", msg), "Enter", "Sleep 250ms",
		"Type \"   sleep 1\"", "Enter", "Sleep 250ms",
		"Type \"done\"", "Enter", "Sleep 250ms",
		"Sleep 6s",
	}

	inputHash := "waiting"

	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

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
