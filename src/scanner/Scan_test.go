package scanner_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

func TestScanDirectory(t *testing.T) {
	b := build.New("testdata")
	files, errors := scanner.Scan(b)
	err := consume(t, files, errors)
	assert.Nil(t, err)
}

func TestScanFile(t *testing.T) {
	b := build.New("testdata/file.q")
	files, errors := scanner.Scan(b)
	err := consume(t, files, errors)
	assert.Nil(t, err)
}

func TestScanNotExisting(t *testing.T) {
	b := build.New("_")
	files, errors := scanner.Scan(b)
	err := consume(t, files, errors)
	assert.NotNil(t, err)
}

func consume(t *testing.T, files <-chan *fs.File, errors <-chan error) error {
	var lastError error

	for files != nil || errors != nil {
		select {
		case file, ok := <-files:
			if !ok {
				files = nil
				continue
			}

			t.Log(file)

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}

			lastError = err
		}
	}

	return lastError
}