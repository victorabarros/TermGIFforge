package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
)

var (
	port = "9001"

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
	rpcGroup.GET("/gif", GetTerminalGift)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Errorf("%+2v/n", err)
	}
}

func GetTerminalGift(c *gin.Context) {
	// TODO check if /output does not exist and create if not
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		return // err
	}
	fmt.Printf("cmdInput %+2v %T \n", cmdInput[0], cmdInput)
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
