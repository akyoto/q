package cli

import "fmt"

// expectedParameterError is created when a command line parameter is missing.
type expectedParameterError struct {
	Parameter string
}

// Error implements the error interface.
func (err *expectedParameterError) Error() string {
	return fmt.Sprintf("Expected parameter '%s'", err.Parameter)
}

// invalidValueError is created when a parameter has an invalid value.
type invalidValueError struct {
	Value     string
	Parameter string
}

// Error implements the error interface.
func (err *invalidValueError) Error() string {
	return fmt.Sprintf("Invalid value '%s' for parameter '%s'", err.Value, err.Parameter)
}

// unknownParameterError is created when a command line parameter is not recognized.
type unknownParameterError struct {
	Parameter string
}

// Error implements the error interface.
func (err *unknownParameterError) Error() string {
	return fmt.Sprintf("Unknown parameter '%s'", err.Parameter)
}