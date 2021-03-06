package errors

import "fmt"

// InvalidType represents an error where a type requirement was not met.
type InvalidType struct {
	Name          string
	Expected      string
	ParameterName string
}

func (err *InvalidType) Error() string {
	if err.ParameterName != "" {
		return fmt.Sprintf("Expected parameter '%s' of type '%s' (encountered '%s')", err.ParameterName, err.Expected, err.Name)
	}

	return fmt.Sprintf("Expected type '%s' instead of '%s'", err.Expected, err.Name)
}
