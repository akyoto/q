package errors

import "fmt"

// UnknownType represents unknown types.
type UnknownType struct {
	Name        string
	CorrectName string
}

func (err *UnknownType) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown type '%s', did you mean '%s'?", err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Unknown type '%s'", err.Name)
}
