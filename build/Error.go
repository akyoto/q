package build

import (
	"fmt"

	"github.com/akyoto/q/token"
)

// Error is a compiler error at a given line and column.
type Error struct {
	Path    string
	Line    int
	Column  int
	Message string
}

// NewError generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func NewError(message string, path string, tokens []token.Token) *Error {
	var (
		lineCount = 0
		lineStart = 0
	)

	for _, oldToken := range tokens {
		if oldToken.Kind == token.NewLine {
			lineCount++
			lineStart = oldToken.Position
		}
	}

	cursorToken := tokens[len(tokens)-1]
	column := cursorToken.Position - lineStart
	return &Error{path, lineCount, column, message}
}

// Error generates the string representation.
func (e *Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.Path, e.Line, e.Column, e.Message)
}
