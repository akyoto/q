package log

import (
	"log"
	"os"
)

var (
	// Info is used for general info messages.
	Info = log.New(os.Stdout, "", 0)

	// Error is used for error messages.
	Error = log.New(os.Stderr, "", 0)
)
