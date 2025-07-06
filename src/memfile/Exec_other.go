//go:build !linux

package memfile

import (
	"os"
	"os/exec"
)

// Exec executes an in-memory file.
func Exec(file *os.File) error {
	err := file.Chmod(0o700)

	if err != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	cmd := exec.Command(file.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}