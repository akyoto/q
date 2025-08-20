package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	MissingInitFunction = errors.String("Missing init function")
	MissingMainFunction = errors.String("Missing main function")
)

// UnknownType error is created when a type could not be found.
type UnknownType struct {
	Name string
}

func (err *UnknownType) Error() string {
	return fmt.Sprintf("Unknown type '%s'", err.Name)
}

// UnusedImport error is created when an import is never used.
type UnusedImport struct {
	Package string
}

func (err *UnusedImport) Error() string {
	return fmt.Sprintf("Unused import '%s'", err.Package)
}