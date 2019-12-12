package build

import (
	"path/filepath"
	"strings"

	"github.com/akyoto/directory"
)

// FindSourceFiles returns all source files in the directory.
func FindSourceFiles(directory string) (<-chan *File, <-chan error) {
	files := make(chan *File)
	errors := make(chan error)
	go findSourceFiles(directory, files, errors)
	return files, errors
}

// findSourceFiles returns all source files in the directory without channel allocations.
func findSourceFiles(dir string, files chan<- *File, errors chan<- error) {
	defer close(files)
	defer close(errors)

	directory.Walk(dir, func(name string) {
		if !strings.HasSuffix(name, ".q") {
			return
		}

		fullPath := filepath.Join(dir, name)
		files <- NewFile(fullPath)
	})
}
