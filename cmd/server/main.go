package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
	"github.com/victorabarros/Terminal-GIFs-API/internal/id"
)

type GIFStatus string

var (
	GIFStatuses = struct {
		Fail       GIFStatus
		Processing GIFStatus
		Ready      GIFStatus
	}{
		Fail:       GIFStatus("Fail"),
		Processing: GIFStatus("Processing"),
		Ready:      GIFStatus("Ready"),
	}
	port = "80"

	setCmds = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
		// TODO add more delay between typing
	}

	// for now, cache is a map inputHash
	cache = map[string]GIFStatus{}
)

func init() {
	// TODO mkdir output/ if not exists
	// TODO list GIFs from output/ and set on "cache"
	// TODO if outGifPath already exist, don't need to redo
	if err := waitingGIF(); err != nil {
		log.Printf("creating waiting: %+2v\n", err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/victorabarros/TermGIFforge?tab=readme-ov-file#termgifforge-")
	})

	rpcGroup := r.Group("/api/v1")
	rpcGroup.GET("/gif", GetTerminalGIF)
	rpcGroup.GET("/mock", GetMockTerminalGIF)

	fmt.Println("Starting app in port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("%+2v/n", err)
	}
}

func GetMockTerminalGIF(c *gin.Context) {
	c.File("output/waiting.gif")
}

func GetTerminalGIF(c *gin.Context) {
	cmdsInputStr := c.Query("commands")
	inputHash := id.NewUUUIDAsString(cmdsInputStr)

	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)
	if status, ok := cache[inputHash]; ok {
		if status == GIFStatuses.Fail {
			// do nothing
		}
		if status == GIFStatuses.Processing {
			waitGifPath := fmt.Sprintf("output/%s.gif", "waiting")
			// TODO check if waiting.gif exists (it takes few seconds when app is starting)
			c.File(waitGifPath)
			return
		}
		if status == GIFStatuses.Ready {
			c.File(outGifPath)
			return
		}
	}

	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Printf("Error trying to serialize object: %+2v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	waitGifPath := fmt.Sprintf("output/%s.gif", "waiting")
	c.File(waitGifPath)

}

func processGIF(id string, cmds []string) error {
	outTapePath := fmt.Sprintf("output/%s.tape", id)
	// TODO introduce mutex here to avoid race condition
	cache[id] = GIFStatuses.Processing

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		// TODO introduce mutex here to avoid race condition
		cache[id] = GIFStatuses.Fail
		return err
	}

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		// TODO introduce mutex here to avoid race condition
		cache[id] = GIFStatuses.Fail
		return err
	}

	// TODO introduce mutex here to avoid race condition
	cache[id] = GIFStatuses.Ready

	exec.Command("rm", "-f", outTapePath).Run()

	return nil
}

func waitingGIF() error {
	// TODO add waiting to "cache"
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
	log.Println(fmt.Sprintf("Type \"   echo \"%s\"\"", "Enter", "Sleep 250ms", msg))

	inputHash := "waiting"

	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	go processGIF(inputHash, cmds)

	return nil
}
