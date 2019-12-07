package errors

import "fmt"

// UnknownFunction represents unknown functions.
type UnknownFunction struct {
	Name        string
	CorrectName string
}

func (err *UnknownFunction) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown function '%s', did you mean '%s'?", err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Unknown function '%s'", err.Name)
}
