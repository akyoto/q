package errors

import "fmt"

// UnknownFunction represents unknown functions.
type UnknownFunction struct {
	FunctionName string
	CorrectName  string
}

func (err *UnknownFunction) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown function '%s', did you mean '%s'?", err.FunctionName, err.CorrectName)
	}

	return fmt.Sprintf("Unknown function '%s'", err.FunctionName)
}
