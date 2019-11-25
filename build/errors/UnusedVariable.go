package errors

import (
	"fmt"
)

// UnusedVariable represents unused variables.
type UnusedVariable struct {
	VariableName string
}

func (err *UnusedVariable) Error() string {
	return fmt.Sprintf("Variable '%s' has never been used", err.VariableName)
}
