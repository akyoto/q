package errors_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestAbsolutePath(t *testing.T) {
	relPath := "../../examples/hello/hello.q"
	absPath, abserr := filepath.Abs(relPath)
	assert.Nil(t, abserr)
	err := test(t, absPath)
	assert.Equal(t, err.Path(), relPath)
}

func TestRelativePath(t *testing.T) {
	relPath := "../../examples/hello/hello.q"
	err := test(t, relPath)
	assert.Equal(t, err.Path(), relPath)
}

func test(t *testing.T, path string) *errors.FileError {
	contents, oserr := os.ReadFile(path)
	assert.Nil(t, oserr)
	tokens := token.Tokenize(contents)

	file := &fs.File{
		Path:   path,
		Bytes:  contents,
		Tokens: tokens,
	}

	main := token.Position(bytes.Index(contents, []byte("main()")))
	err := errors.New(io.EOF, file, main)
	assert.NotNil(t, err)

	line, column := err.LineColumn()
	assert.Equal(t, line, 3)
	assert.Equal(t, column, 1)
	assert.Equal(t, err.Unwrap().Error(), "EOF")
	assert.Contains(t, err.Error(), ":3:1: EOF")
	return err
}