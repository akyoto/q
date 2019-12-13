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
	return fmt.Sprintf("%v\n\n%s", err.Err, log.Faint.Sprint(err.Stack))
}

// New creates a new error with stack information.
func New(err error) *WithStack {
	buffer := make([]byte, 4096)
	n := runtime.Stack(buffer, false)
	stack := string(buffer[:n])
	lines := strings.Split(stack, "\n")
	var extractedLines []string

	for i := 4; i < len(lines); i += 2 {
		line := strings.TrimSpace(lines[i])
		space := strings.LastIndex(line, " ")

		if space != -1 {
			line = line[:space]
		}

		extractedLines = append(extractedLines, line)
	}

	stack = strings.Join(extractedLines, "\n")

	return &WithStack{
		Err:   err,
		Stack: stack,
	}
}
