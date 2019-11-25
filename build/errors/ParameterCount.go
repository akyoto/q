package errors

import "fmt"

// ParameterCount represents an error where the parameter count is different from the expected number of parameters.
type ParameterCount struct {
	FunctionName  string
	CountGiven    int
	CountRequired int
}

func (err *ParameterCount) Error() string {
	if err.CountGiven < err.CountRequired {
		return fmt.Sprintf("Too few arguments in '%s' call", err.FunctionName)
	}

	if err.CountGiven > err.CountRequired {
		return fmt.Sprintf("Too many arguments in '%s' call", err.FunctionName)
	}

	return ""
}
