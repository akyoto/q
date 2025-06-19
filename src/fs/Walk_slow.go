//go:build !linux && !darwin

package fs

import "os"

// Walk calls your callback function for every file name inside the directory.
// It doesn't distinguish between files and directories.
func Walk(directory string, callBack func(string)) error {
	f, err := os.Open(directory)

	if err != nil {
		return err
	}

	files, err := f.Readdirnames(0)
	f.Close()

	if err != nil {
		return err
	}

	for _, file := range files {
		callBack(file)
	}

	return nil
}