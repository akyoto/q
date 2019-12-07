package errors

import "fmt"

// UnknownVariable represents unknown variables.
type UnknownVariable struct {
	Name string
}

func (err *UnknownVariable) Error() string {
	return fmt.Sprintf("Unknown variable: '%s'", err.Name)
}
