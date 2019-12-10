package build

import (
	"os"
	"path/filepath"
	"strings"
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

	// Fix stdLib path for tests and benchmarks
	if err != nil {
		qRoot, err = os.Getwd()

		if err != nil {
			return "", err
		}

		// A little hacky, but it works
		qPos := strings.LastIndex(qRoot, "/q")
		qRoot = qRoot[:qPos+2]

		return filepath.Join(qRoot, "lib"), nil
	}

	return stdLib, nil
}
