package global

import (
	"fmt"
	"os"
	"path/filepath"
)

// findLibrary tries to go up each directory from the working directory and check for the existence of a "lib" directory.
// This is needed for tests to work correctly.
func findLibrary() {
	dir := WorkingDirectory

	for {
		Library = filepath.Join(dir, "lib")
		stat, err := os.Stat(Library)

		if err == nil && stat.IsDir() {
			return
		}

		newDir := filepath.Dir(dir)

		if newDir == dir {
			fmt.Fprintln(os.Stderr, "standard library not found")
			os.Exit(1)
		}

		dir = newDir
	}
}