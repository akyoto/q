package build

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// FindStandardLibrary returns the path to the standard library.
func FindStandardLibrary() (string, error) {
	compiler, err := os.Executable()

	if err != nil {
		return "", err
	}

	qRoot := filepath.Dir(compiler)
	stdLib := filepath.Join(qRoot, "lib")
	_, err = os.Stat(stdLib)

	if !os.IsNotExist(err) {
		return stdLib, err
	}

	// Go up from current directory until we find a lib directory
	qRoot, err = os.Getwd()

	if err != nil {
		return "", err
	}

	for {
		stdLib = path.Join(qRoot, "lib")
		stat, err := os.Stat(stdLib)

		if !os.IsNotExist(err) && stat != nil && stat.IsDir() {
			return stdLib, err
		}

		if qRoot == "/" {
			return "", errors.New("Standard library not found")
		}

		// Go up one level
		qRoot = path.Clean(path.Join(qRoot, ".."))
	}
}
