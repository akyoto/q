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

	go func() {
		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if path == directory {
				return nil
			}

			if info.IsDir() {
				return filepath.SkipDir
			}

			if !strings.HasSuffix(path, ".q") {
				return nil
			}

			files <- NewFile(path)
			return nil
		})

		if err != nil {
			errors <- err
		}

		close(files)
		close(errors)
	}()

	return files, errors
}
