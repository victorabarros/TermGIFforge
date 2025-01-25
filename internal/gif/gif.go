package gif

import (
	"os"
	"os/exec"
)

func WriteTape(cmds []string, filePath string) error {
	// Remove old file
	if err := exec.Command("rm", "-f", filePath).Run(); err != nil {
		return err
	}

	// Create or overwrite the file
	file, err := os.Create(filePath)
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

func ExecVHS(filePath string) error {
	cmd := exec.Command("vhs", filePath)

	// Set the output to the current process's stdout and stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
