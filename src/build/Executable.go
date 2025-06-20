package build

import (
	"path/filepath"
	"strings"
)

// Executable returns the path to the executable.
func (build *Build) Executable() string {
	path, err := filepath.Abs(build.Files[0])

	if err != nil {
		panic(err)
	}

	if strings.HasSuffix(path, ".q") {
		path = fromFileName(path)
	} else {
		path = fromDirectoryName(path)
	}

	if build.OS == Windows {
		path += ".exe"
	}

	return path
}

// fromFileName returns the executable path based on the file name.
func fromFileName(path string) string {
	return filepath.Join(filepath.Dir(path), strings.TrimSuffix(filepath.Base(path), ".q"))
}

// fromDirectoryName returns the executable path based on the directory name.
func fromDirectoryName(path string) string {
	return filepath.Join(path, filepath.Base(path))
}