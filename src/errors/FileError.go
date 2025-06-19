package errors

import (
	"fmt"
	"os"
	"path/filepath"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// FileError is an error inside a file at a given line and column.
type FileError struct {
	err      error
	file     *fs.File
	position token.Position
	stack    string
}

// Error implements the error interface.
func (e *FileError) Error() string {
	path := e.Path()
	line, column := e.LineColumn()
	return fmt.Sprintf("%s:%d:%d: %s\n\n%s", path, line, column, e.err, e.stack)
}

// LineColumn returns the line and column of the error.
func (e *FileError) LineColumn() (int, int) {
	line := 1
	column := 1
	lineStart := -1

	for _, t := range e.file.Tokens {
		if t.Position >= e.position {
			column = int(e.position) - lineStart
			break
		}

		if t.Kind == token.NewLine {
			lineStart = int(t.Position)
			line++
		}
	}

	return line, column
}

// Path returns the relative path of the file to shorten the error message.
func (e *FileError) Path() string {
	cwd, err := os.Getwd()

	if err != nil {
		return e.file.Path
	}

	relative, err := filepath.Rel(cwd, e.file.Path)

	if err != nil {
		return e.file.Path
	}

	return relative
}

// Unwrap returns the wrapped error.
func (e *FileError) Unwrap() error {
	return e.err
}