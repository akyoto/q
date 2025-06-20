package global

import (
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

		if dir == "/" {
			panic("standard library not found")
		}

		dir = filepath.Dir(dir)
	}
}