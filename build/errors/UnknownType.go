package errors

import "fmt"

// UnknownType represents unknown types.
type UnknownType struct {
	Name string
}

func (err *UnknownType) Error() string {
	return fmt.Sprintf("Unknown type: '%s'", err.Name)
}
