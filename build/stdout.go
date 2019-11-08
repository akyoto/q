package build

import (
	"log"
	"os"
)

var (
	// stdout is used instead of os.Stdout for goroutine-safe logging.
	stdout = log.New(os.Stdout, "", 0)

	// stderr is used instead of os.Stderr for goroutine-safe logging.
	stderr = log.New(os.Stderr, "", 0)
)
