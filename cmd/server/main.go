package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
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
	if err := waitingGIF(); err != nil {
		log.Printf("creating waiting: %+2v\n", err)
	}
}

func waitingGIF() error {
	cmdInput := []string{
		// TODO increase font
		"Set FontSize 15",
		"Type \"while true; do\"", "Sleep 200ms", "Enter", "Sleep 200ms",
		"Type \"   echo \"Wait...\"\"", "Sleep 200ms", "Enter", "Sleep 200ms",
		"Type \"   sleep 1\"", "Sleep 200ms", "Enter", "Sleep 200ms",
		"Type \"done\"", "Sleep 200ms", "Enter", "Sleep 200ms",
		"Sleep 6s",
	}

	inputHash := "waiting"

	outTapePath := fmt.Sprintf("output/%s.tape", inputHash)
	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	// TODO introduce mutex here to avoid race condition
	cache[inputHash] = GIFStatuses.Processing

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		// TODO introduce mutex here to avoid race condition
		cache[inputHash] = GIFStatuses.Fail
		return err
	}

	if err := gif.ExecVHS(outTapePath); err != nil {
		// TODO introduce mutex here to avoid race condition
		cache[inputHash] = GIFStatuses.Fail
		return err
	}

	// TODO introduce mutex here to avoid race condition
	cache[inputHash] = GIFStatuses.Ready

	exec.Command("rm", "-f", outTapePath).Run()
	return nil
}

func main() {
	r := gin.Default()

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
	inputHash := newUUUID(cmdsInputStr)

	outTapePath := fmt.Sprintf("output/%s.tape", inputHash)
	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)
	if status, ok := cache[inputHash]; ok {
		if status == GIFStatuses.Fail {
			// do nothing
		}
		if status == GIFStatuses.Processing {
			// TODO use a GIF of a loop echoing "WAIT"
			waitGifPath := fmt.Sprintf("output/%s.gif", inputHash)
			c.File(waitGifPath)
			// c.JSON(http.StatusAccepted, gin.H{"message": "wait"})
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

	// TODO introduce mutex here to avoid race condition
	cache[inputHash] = GIFStatuses.Processing

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		// TODO introduce mutex here to avoid race condition
		cache[inputHash] = GIFStatuses.Fail
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		// TODO introduce mutex here to avoid race condition
		cache[inputHash] = GIFStatuses.Fail
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// exec.Command("mv", "demo.gif", "output/", gif.FileName).Run()
	// TODO introduce mutex here to avoid race condition
	cache[inputHash] = GIFStatuses.Ready
	c.File(outGifPath)

	exec.Command("rm", "-f", outTapePath).Run()
}

// create deterministic UUUID
func newUUUID(input string) string {
	// calculate the MD5 hash of the
	md5hash := md5.New()
	_, err := md5hash.Write([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	// convert the hash value to a string
	md5string := hex.EncodeToString(md5hash.Sum(nil))

	// generate the UUID from the
	uuid, err := uuid.FromBytes([]byte(md5string[0:16]))
	if err != nil {
		log.Fatal(err)
	}

	return uuid.String()
}
