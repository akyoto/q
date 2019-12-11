package build

import (
	"os"
	"path/filepath"
	"strings"
)

// FindSourceFiles returns all source files in the directory.
func FindSourceFiles(directory string) (<-chan *File, <-chan error) {
	files := make(chan *File)
	errors := make(chan error)
	go findSourceFiles(directory, files, errors)
	return files, errors
}

// findSourceFiles returns all source files in the directory without channel allocations.
func findSourceFiles(directory string, files chan<- *File, errors chan<- error) {
	defer close(files)
	defer close(errors)

	fd, err := os.Open(directory)

	if err != nil {
		errors <- err
		return
	}

	defer fd.Close()

	names, err := fd.Readdirnames(0)

	if err != nil {
		errors <- err
		return
	}

	for _, name := range names {
		if !strings.HasSuffix(name, ".q") {
			continue
		}

		fullPath := filepath.Join(directory, name)
		files <- NewFile(fullPath)
	}
}
