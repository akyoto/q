//go:build linux

package memfile

import (
	"fmt"
	"os"
	"os/exec"
)

// Exec executes an in-memory file.
func Exec(file *os.File) error {
	defer file.Close()
	cmd := exec.Command(fmt.Sprintf("/proc/self/fd/%d", file.Fd()))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}