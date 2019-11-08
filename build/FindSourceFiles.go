package build

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/akyoto/q/build/log"
)

// FindSourceFiles returns all source files in the directory (top-level only, not recursive).
func FindSourceFiles(directory string) <-chan *File {
	files := make(chan *File)

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
			log.Error.Println(err)
			os.Exit(1)
		}

		close(files)
	}()

	return files
}
