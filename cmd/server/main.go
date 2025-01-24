package main

import (
	"fmt"

	"github.com/victorabarros/Terminal-GIFs-API/internal/gif"
)

func main() {
	cmds := []string{
		"Output demo.gif",
		"Set FontSize 26",
		"Set Width 1200",
		"Set Height 600",
		"Type \"echo 'Welcome to VHS!'\"",
		"Sleep 100ms",
		"Enter",
		"Sleep 100ms",
		"Type \"ls -a\"",
		"Sleep 100ms",
		"Enter",
		"Sleep 1s",
	}

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
}
