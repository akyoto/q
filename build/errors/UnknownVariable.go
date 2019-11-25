package errors

import "fmt"

// UnknownVariable represents unknown variables.
type UnknownVariable struct {
	VariableName string
}

func (err *UnknownVariable) Error() string {
	return fmt.Sprintf("Unknown variable: '%s'", err.VariableName)
}
