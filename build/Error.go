package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Error is a compiler error at a given line and column.
type Error struct {
	Path     string
	Line     int
	Column   int
	Function *Function
	Err      error
}

// NewError generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func NewError(err error, path string, tokens []token.Token, function *Function) *Error {
	var (
		lineCount = 0
		lineStart = 0
	)

	for i, oldToken := range tokens {
		if oldToken.Kind == token.NewLine && i != len(tokens)-1 {
			lineCount++
			lineStart = oldToken.Position
		}
	}

	cursorToken := tokens[len(tokens)-1]
	column := cursorToken.Position - lineStart

	cursorRight, ok := err.(errors.CursorRight)

	if ok && cursorRight.CursorRight() {
		column += len(cursorToken.Bytes)
	}

	return &Error{path, lineCount, column, function, err}
}

// Error generates the string representation.
func (e *Error) Error() string {
	path := e.Path
	cwd, err := os.Getwd()

	if err == nil {
		relativePath, err := filepath.Rel(cwd, e.Path)

		if err == nil {
			path = relativePath
		}
	}

	if e.Function != nil {
		return fmt.Sprintf("%s:%d:%d: [%s] %s", path, e.Line, e.Column, e.Function.Name, e.Err)
	}

	return fmt.Sprintf("%s:%d:%d: %s", path, e.Line, e.Column, e.Err)
}
