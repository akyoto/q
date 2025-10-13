package errors

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// FileError is an error inside a file at a given line and column.
type FileError struct {
	err    error
	file   *fs.File
	stack  string
	source Source
}

// Error implements the error interface.
func (e *FileError) Error() string {
	return e.err.Error()
}

// File returns the file the error came from.
func (e *FileError) File() *fs.File {
	return e.file
}

// Line returns the line contents and the offset from the position.
func (e *FileError) Line() (string, int) {
	contents := e.file.Bytes
	start := bytes.LastIndexByte(contents[:e.source.Start()], '\n')

	if start == -1 {
		start = 0
	} else {
		start++
	}

	end := bytes.IndexByte(contents[e.source.Start():], '\n')

	if end == -1 {
		end = len(contents)
	} else {
		end += int(e.source.Start())
	}

	asciiSpace := [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

	for asciiSpace[contents[start]] == 1 {
		start++
	}

	return string(contents[start:end]), int(e.source.Start()) - start
}

// LineColumn returns the line and column of the error.
func (e *FileError) LineColumn() (int, int) {
	line := 1
	column := 1
	lineStart := -1

	for _, t := range e.file.Tokens {
		if t.Position >= e.source.Start() {
			column = int(e.source.Start()) - lineStart
			break
		}

		if t.Kind == token.NewLine {
			lineStart = int(t.Position)
			line++
		}
	}

	return line, column
}

// Link returns a clickable link containing the path, line and column.
func (e *FileError) Link() string {
	path := e.Path()
	line, column := e.LineColumn()
	return fmt.Sprintf("%s:%d:%d", path, line, column)
}

// Path returns the relative path of the file to shorten the error message.
func (e *FileError) Path() string {
	workDir, err := os.Getwd()

	if err != nil {
		return e.file.Path
	}

	relative, err := filepath.Rel(workDir, e.file.Path)

	if err != nil {
		return e.file.Path
	}

	return filepath.ToSlash(relative)
}

// Source returns the byte offsets inside the file.
func (e *FileError) Source() Source {
	return e.source
}

// Stack returns the call stack at the time the error was created.
func (e *FileError) Stack() string {
	return e.stack
}

// Unwrap returns the wrapped error.
func (e *FileError) Unwrap() error {
	return e.err
}