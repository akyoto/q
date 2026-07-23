package compiler

import (
	"fmt"
	"strings"

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
	tmp := strings.Builder{}

	for _, single := range err.Errors {
		tmp.WriteString(single.Error())
		tmp.WriteByte('\n')
	}

	return tmp.String()
}

// UnusedImport error is created when an import is never used.
type UnusedImport struct {
	Package string
}

func (err *UnusedImport) Error() string {
	return fmt.Sprintf("Unused import '%s'", err.Package)
}