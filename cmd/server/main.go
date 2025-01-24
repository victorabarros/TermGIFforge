package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	cmdsInputStr := c.Query("commands")
	cmdInput := []string{}
	if err := json.Unmarshal([]byte(cmdsInputStr), &cmdInput); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		return // err
	}
	fmt.Printf("cmdInput %+2v %T \n", cmdInput, cmdInput)
	cmds = append(cmds, cmdInput...)
	fmt.Printf("cmdInput %+2v %T \n", cmds, cmds)

	if err := gif.WriteTape(cmds); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	if err := gif.ExecVHS(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		return
	}

	// exec.Command("mv", "demo.gif", "output/", gif.FileName).Run()
	// exec.Command("rm", "-f", gif.FileName).Run()

	fmt.Println("Command executed successfully!")
	c.File(gif.GifFileName)
}
