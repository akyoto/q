package errors

import "fmt"

// InvalidTopLevel error is created when a top-level instruction is not valid.
type InvalidTopLevel struct {
	Instruction string
}

// Error implements the error interface.
func (err *InvalidTopLevel) Error() string {
	return fmt.Sprintf("Invalid top level instruction '%s'", err.Instruction)
}