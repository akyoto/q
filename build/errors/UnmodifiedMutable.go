package errors

import "fmt"

// UnmodifiedMutable represents mutable variables that are not modified.
type UnmodifiedMutable struct {
	VariableName string
}

func (err *UnmodifiedMutable) Error() string {
	return fmt.Sprintf("Mutable variable '%s' has never been modified", err.VariableName)
}
