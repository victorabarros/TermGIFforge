package gif

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	FileName    = "output/demo.tape"
	GifFileName = "output/demo.gif"
)

func ExecVHS() error {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Prepare the docker command
	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/vhs", workingDir),
		"ghcr.io/charmbracelet/vhs", FileName,
	)

	// Set the output to the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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

	fmt.Printf("Commands written to %s successfully!\n", FileName)
	return nil
}
