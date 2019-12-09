package errors

import "fmt"

// MissingType represents an error where a type is missing.
type MissingType struct {
	Of string
}

func (err *MissingType) Error() string {
	return fmt.Sprintf("Missing type of '%s'", err.Of)
}
