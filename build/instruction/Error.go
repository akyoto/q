package instruction

import (
	"github.com/akyoto/q/build/token"
)

type Error struct {
	Message         string
	Position        token.Position
	RightSideCursor bool
}

func (err *Error) Error() string {
	return err.Message
}

func (err *Error) CursorRight() bool {
	return err.RightSideCursor
}
