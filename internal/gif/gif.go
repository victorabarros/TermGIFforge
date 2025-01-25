package gif

import (
	"os"
	"os/exec"
)

const (
	FileName    = "output/demo.tape"
	GifFileName = "output/demo.gif"
)

func ExecVHS() error {
	cmd := exec.Command("vhs", FileName)

	// Set the output to the current process's stdout and stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func WriteTape(cmds []string) error {

	// Create or overwrite the file
	file, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Write each command to the file
	for _, cmd := range cmds {
		_, err = file.WriteString(cmd + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
