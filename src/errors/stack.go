package errors

import (
	"runtime"
	"strings"
)

// stack generates a stack trace.
func stack() string {
	var (
		final  []string
		buffer = make([]byte, 4096)
		n      = runtime.Stack(buffer, false)
		stack  = string(buffer[:n])
		lines  = strings.Split(stack, "\n")
		limit  = min(len(lines), 16)
	)

	for i := 6; i < limit; i += 2 {
		line := strings.TrimSpace(lines[i])
		space := strings.LastIndex(line, " ")

		if space != -1 {
			line = line[:space]
		}

		final = append(final, line)
	}

	return strings.Join(final, "\n")
}