package errors

import "fmt"

// MissingReturnValue represents an error where an empty return statement
// was encountered in a function that is supposed to return a value.
type MissingReturnValue struct {
	ReturnType string
}

func (err *MissingReturnValue) Error() string {
	return fmt.Sprintf("Missing return value of type '%s'", err.ReturnType)
}
