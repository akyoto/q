package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	MissingInitFunction = errors.String("Missing init function")
	MissingMainFunction = errors.String("Missing main function")
)

// MultiError error is created when there is more than one error.
type MultiError struct {
	Errors []error
}

func (err *MultiError) Error() string {
	return ""
}

// UnusedImport error is created when an import is never used.
type UnusedImport struct {
	Package string
}

func (err *UnusedImport) Error() string {
	return fmt.Sprintf("Unused import '%s'", err.Package)
}