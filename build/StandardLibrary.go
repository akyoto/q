package build

import (
	"os"
	"path/filepath"
)

// stdLibPath returns the path to the standard library.
func stdLibPath() (string, error) {
	compiler, err := os.Executable()

	if err != nil {
		return "", err
	}

	qRoot := filepath.Dir(compiler)
	stdLib := filepath.Join(qRoot, "lib")
	_, err = os.Stat(stdLib)

	// Fix stdLib path for tests inside the "build" directory
	if err != nil {
		qRoot, err = os.Getwd()

		if err != nil {
			return "", err
		}

		return filepath.Join(qRoot, "..", "lib"), nil
	}

	return stdLib, nil
}
