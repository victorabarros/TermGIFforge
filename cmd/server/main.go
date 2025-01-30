package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/termgifforge/internal/files"
	"github.com/victorabarros/termgifforge/internal/gif"
	"github.com/victorabarros/termgifforge/internal/id"
	"github.com/victorabarros/termgifforge/pkg/models"
)

var (
	port    = "80"
	version = "0.1.1"

	outputCmdFormat = "Output %s"
	setCmds         = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	details = models.NewGIFDetails()
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
		details.SetStatus(id, models.GIFStatuses.Ready)
	}

	if d, _ := details.Get("error"); d.Status != models.GIFStatuses.Ready {
		errorGIF()
	}
	if d, _ := details.Get("invalid"); d.Status != models.GIFStatuses.Ready {
		invalidGIF()
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(
			http.StatusTemporaryRedirect,
			"https://victor.barros.engineer/termgif",
		)
	})

	r.GET("/ping", func(c *gin.Context) {
		gifs, err := files.ListGIFs()
		if err != nil {
			log.Printf("Fail to list GIFs %+2v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "GIF in process"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "nofGIFs": len(gifs), "version": version})
	})

	rpcGroup := r.Group("/api/v1")
	rpcGroup.GET("/gif", terminalGIF)
	rpcGroup.GET("/mock", func(c *gin.Context) {
		c.File("output/error.gif")
	})
	rpcGroup.GET("/internal/clean/:id", func(c *gin.Context) {
		//TODO user files.EraseGIF
	})

	go files.Cleaner(&details)

	log.Printf("Starting app version %s in port %s", version, port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Printf("%+2v/n", err)
	}
}

func terminalGIF(c *gin.Context) {
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Printf("Error trying to serialize object: %+2v\n", err)
		c.File("output/invalid.gif")
		return
	}

	id := id.NewUUUIDAsString(cmdsInputStr)
	outGifPath := fmt.Sprintf("output/%s.gif", id)
	if d, ok := details.Get(id); ok {
		if d.Status == models.GIFStatuses.Fail {
			c.File("output/error.gif")
			return
		}
		if d.Status == models.GIFStatuses.Processing {
			c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
			return
		}
		if d.Status == models.GIFStatuses.Ready {
			details.SetLastAccess(id, time.Now())
			c.File(outGifPath)
			return
		}
	}

	cmds := append([]string{fmt.Sprintf(outputCmdFormat, outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	c.JSON(http.StatusAccepted, gin.H{"message": "GIF in process"})
}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	details.SetStatus(id, models.GIFStatuses.Processing)

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		details.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}
	defer os.Remove(outTapePath)

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		details.SetStatus(id, models.GIFStatuses.Fail)
		return err
	}

	details.SetStatus(id, models.GIFStatuses.Ready)

	log.Printf("GIF Created id %s\n", id)
	return nil
}

func errorGIF() error {
	cmdInput := []string{
		"Type \"Sorry, it was not possible create your GIF. =/\"",
		"Sleep 6s",
	}

	id := "error"
	outGifPath := fmt.Sprintf("output/%s.gif", id)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	return nil
}

func invalidGIF() error {
	cmdInput := []string{
		"Type \"Invalid request...\"",
		"Sleep 6s",
	}

	id := "invalid"
	outGifPath := fmt.Sprintf("output/%s.gif", id)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(id, cmds)

	return nil
}
