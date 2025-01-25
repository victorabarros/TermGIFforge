package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
)

var (
	port = "80"

	cmds = []string{
		fmt.Sprintf("Output %s", gif.GifFileName),
		"Set WindowBar Colorful",
		"Set FontSize 15",
		"Set Width 600",
		"Set Height 300",
	}
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
	c.File(gif.GifFileName)
}

func GetTerminalGIF(c *gin.Context) {
	// ctx := c.Request.Context()

	getTerminalGIF(c)
}

func getTerminalGIF(c *gin.Context) {
	// TODO check if /output does not exist and create if not
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		log.Println("Error running command: %v\n", err)
		return // err
	}
	log.Println("cmdInput %+2v \n", cmdInput[0])
	cmds = append(cmds, cmdInput...)

	func() {
		mt.Lock()
		if err := gif.WriteTape(cmds); err != nil {
			log.Println("Error writing to file: %v\n", err)
			return
		}

		if err := gif.ExecVHS(); err != nil {
			log.Println("Error running command: %v\n", err)
			return
		}
		mt.Unlock()
	}()

	// exec.Command("mv", "demo.gif", "output/", gif.FileName).Run()

	c.File(gif.GifFileName)

	exec.Command("rm", "-f", gif.GifFileName).Run()
	exec.Command("rm", "-f", gif.FileName).Run()
}
