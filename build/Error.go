package build

import (
	"fmt"
)

type Error struct {
	Path    string
	Line    int
	Column  int
	Cursor  int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.Path, e.Line, e.Column, e.Message)
}
