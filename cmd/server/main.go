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
	port = "80"

	outputCmdFormat = "Output %s"
	setCmds         = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	statuses   = models.StatusDetails{}
	lastAccess = map[string]time.Time{}
)

func init() {
	statuses = models.NewStatusDetails()

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
		statuses.Set(id, models.GIFStatuses.Ready)
	}

	if status, _ := statuses.Get("error"); status != models.GIFStatuses.Ready {
		errorGIF()
	}
	if status, _ := statuses.Get("invalid"); status != models.GIFStatuses.Ready {
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
	if status, ok := statuses.Get(inputHash); ok {
		if status == models.GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if status == models.GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if status == models.GIFStatuses.Ready {
			// TODO mutex
			lastAccess[inputHash] = time.Now()
			c.File(outGifPath)
			return
		}
	}

	cmds := append([]string{fmt.Sprintf(outputCmdFormat, outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	statuses.Set(id, models.GIFStatuses.Processing)

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		statuses.Set(id, models.GIFStatuses.Fail)
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		statuses.Set(id, models.GIFStatuses.Fail)
		return err
	}

	statuses.Set(id, models.GIFStatuses.Ready)
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
