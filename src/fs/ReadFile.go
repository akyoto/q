package fs

import (
	"os"
)

// ReadFile reads the contents of the given file path into a new buffer.
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	stat, err := f.Stat()

	if err != nil {
		return nil, err
	}

	contents := make([]byte, stat.Size())
	pos := 0

	for {
		n, err := f.Read(contents[pos:])
		pos += n

		if pos >= len(contents) {
			return contents, nil
		}

		if err != nil {
			return nil, err
		}
	}
}