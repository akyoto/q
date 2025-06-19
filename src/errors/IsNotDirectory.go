package errors

import "fmt"

// IsNotDirectory error is created when a path is not a directory.
type IsNotDirectory struct {
	Path string
}

// Error implements the error interface.
func (err *IsNotDirectory) Error() string {
	return fmt.Sprintf("'%s' is not a directory", err.Path)
}