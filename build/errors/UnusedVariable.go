package errors

import (
	"fmt"
)

// UnusedVariable represents unused variables.
type UnusedVariable struct {
	Name string
}

func (err *UnusedVariable) Error() string {
	return fmt.Sprintf("Variable '%s' has never been used", err.Name)
}
