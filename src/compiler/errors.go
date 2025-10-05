package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	MissingInitFunction = errors.String("Missing init function")
	MissingMainFunction = errors.String("Missing main function")
)

// CycleDetected error is created when an invalid dependency cycle was detected.
type CycleDetected struct {
	A string
	B string
}

func (err *CycleDetected) Error() string {
	return fmt.Sprintf("Cycle detected: '%s' depends on '%s' which depends on '%s'", err.A, err.B, err.A)
}

// UnusedImport error is created when an import is never used.
type UnusedImport struct {
	Package string
}

func (err *UnusedImport) Error() string {
	return fmt.Sprintf("Unused import '%s'", err.Package)
}