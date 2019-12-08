package errors

import "fmt"

// UnmodifiedMutable represents mutable variables that are not modified.
type UnmodifiedMutable struct {
	Name string
}

func (err *UnmodifiedMutable) Error() string {
	return fmt.Sprintf("Mutable variable '%s' has never been modified", err.Name)
}
