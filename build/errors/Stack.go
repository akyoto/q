package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/akyoto/q/build/log"
)

// WithStack describes an error with a runtime stack.
type WithStack struct {
	Err   error
	Stack string
}

func (err *WithStack) Error() string {
	return fmt.Sprintf("%v\n%s", err.Err, log.Faint.Sprint(err.Stack))
}

// New creates a new error with stack information.
func New(err error) *WithStack {
	buffer := make([]byte, 4096)
	n := runtime.Stack(buffer, false)
	stack := string(buffer[:n])
	lines := strings.Split(stack, "\n")
	stack = strings.TrimSpace(lines[4])
	space := strings.LastIndex(stack, " ")

	if space != -1 {
		stack = stack[:space]
	}
	// lines = append(lines[:1], lines[3:]...)

	return &WithStack{
		Err:   err,
		Stack: stack,
	}
}
