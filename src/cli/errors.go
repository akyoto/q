package cli

import "fmt"

// ExpectedParameter is created when a command line parameter is missing.
type ExpectedParameter struct {
	Parameter string
}

func (err *ExpectedParameter) Error() string {
	return fmt.Sprintf("Expected parameter '%s'", err.Parameter)
}

// InvalidValue is created when a parameter has an invalid value.
type InvalidValue struct {
	Value     string
	Parameter string
}

func (err *InvalidValue) Error() string {
	return fmt.Sprintf("Invalid value '%s' for parameter '%s'", err.Value, err.Parameter)
}

// UnknownParameter is created when a command line parameter is not recognized.
type UnknownParameter struct {
	Parameter string
}

func (err *UnknownParameter) Error() string {
	return fmt.Sprintf("Unknown parameter '%s'", err.Parameter)
}