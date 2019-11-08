package log

import (
	"log"
	"os"
)

var (
	// Info is used instead of os.Stdout for goroutine-safe logging.
	Info = log.New(os.Stdout, "", 0)

	// Error is used instead of os.Stderr for goroutine-safe logging.
	Error = log.New(os.Stderr, "", 0)
)
