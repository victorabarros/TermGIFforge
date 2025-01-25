package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
)

var (
	port = "80"

	setCmds = []string{
		"Set WindowBar Colorful",
		"Set FontSize 12",
		"Set Width 800",
		"Set Height 400",
	}

	// for now, DB is a map inputHash
	database = map[string]struct{}{}

	mt = sync.Mutex{}
)

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
	c.File("output/demo.gif")
}

func GetTerminalGIF(c *gin.Context) {
	cmdsInputStr := c.Query("commands")
	inputHash := newUUUID(cmdsInputStr)

	outTapePath := fmt.Sprintf("output/%s.tape", inputHash)
	outGifPath := fmt.Sprintf("output/%s.gif", inputHash)
	if _, ok := database[inputHash]; ok {
		c.File(outGifPath)
		return
	}

	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Printf("Error trying to serialize object: %+2v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	cmds := append([]string{fmt.Sprintf("Output %s", outGifPath)}, setCmds...)
	cmds = append(cmds, cmdInput...)

	mt.Lock()
	defer mt.Unlock()

	if err := gif.WriteTape(cmds, outTapePath); err != nil {
		log.Printf("Error writing to file: %+2v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := gif.ExecVHS(outTapePath); err != nil {
		log.Printf("Error running command: %+2v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// exec.Command("mv", "demo.gif", "output/", gif.FileName).Run()
	database[inputHash] = struct{}{}
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
