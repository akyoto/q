package errors

import "fmt"

// UnknownVariable represents unknown variables.
type UnknownVariable struct {
	Name        string
	CorrectName string
}

func (err *UnknownVariable) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown variable '%s', did you mean '%s'?", err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Unknown variable '%s'", err.Name)
}
