package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
)

var (
	port = "80"

	cmds = []string{
		fmt.Sprintf("Output %s", gif.GifFileName),
		"Set FontSize 26",
		"Set Width 1200",
		"Set Height 600",
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
		fmt.Errorf("%+2v/n", err)
	}
}

func GetMockTerminalGIF(c *gin.Context) {
	c.File(gif.GifFileName)
}

func GetTerminalGIF(c *gin.Context) {
	ctx := c.Request.Context()

	// Simulate a long-running process
	done := make(chan bool)
	go func() {
		defer close(done)
		getTerminalGIF(c)
		log.Println("shouldnt show if canceled")
	}()

	select {
	case <-ctx.Done(): // If client cancels the request
		log.Println("Request canceled by the client")
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request canceled by client"})
		return
	case <-done: // If the process completes normally
		// c.JSON(http.StatusOK, gin.H{"status": "process completed successfully"})
		return
	}

}

func getTerminalGIF(c *gin.Context) {
	// TODO check if /output does not exist and create if not
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		return // err
	}
	fmt.Printf("cmdInput %+2v \n", cmdInput[0])
	cmds = append(cmds, cmdInput...)

	func() {
		mt.Lock()
		if err := gif.WriteTape(cmds); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}

		if err := gif.ExecVHS(); err != nil {
			fmt.Printf("Error running command: %v\n", err)
			return
		}
		mt.Unlock()
	}()

	// exec.Command("mv", "demo.gif", "output/", gif.FileName).Run()

	c.File(gif.GifFileName)

	// exec.Command("rm", "-f", gif.GifFileName).Run()
}
