package scanner_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

func TestScanDirectory(t *testing.T) {
	b := build.New("testdata")
	functions, files, errors := scanner.Scan(b)
	err := consume(t, functions, files, errors)
	assert.Nil(t, err)
}

func TestScanFile(t *testing.T) {
	b := build.New("testdata/file.q")
	functions, files, errors := scanner.Scan(b)
	err := consume(t, functions, files, errors)
	assert.Nil(t, err)
}

func TestScanNotExisting(t *testing.T) {
	b := build.New("_")
	functions, files, errors := scanner.Scan(b)
	err := consume(t, functions, files, errors)
	assert.NotNil(t, err)
}

func TestScanHelloExample(t *testing.T) {
	b := build.New("../../examples/hello")
	functions, files, errors := scanner.Scan(b)
	err := consume(t, functions, files, errors)
	assert.Nil(t, err)
}

func consume(t *testing.T, functions <-chan *core.Function, files <-chan *fs.File, errors <-chan error) error {
	var lastError error

	for functions != nil || files != nil || errors != nil {
		select {
		case function, ok := <-functions:
			if !ok {
				functions = nil
				continue
			}

			t.Log(function)

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