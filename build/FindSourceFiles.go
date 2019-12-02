package build

import (
	"os"
	"path/filepath"
	"strings"
)

// FindSourceFiles returns all source files in the directory (top-level only, not recursive).
func FindSourceFiles(directory string) (<-chan *File, <-chan error) {
	files := make(chan *File)
	errors := make(chan error)
	go findSourceFiles(directory, files, errors)
	return files, errors
}

// findSourceFiles returns all source files in the directory (top-level only, not recursive).
func findSourceFiles(directory string, files chan<- *File, errors chan<- error) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if path == directory || !strings.HasSuffix(path, ".q") {
			return nil
		}

		if info.IsDir() {
			return filepath.SkipDir
		}

		files <- NewFile(path)
		return nil
	})

	if err != nil {
		errors <- err
	}

	close(files)
	close(errors)
}
