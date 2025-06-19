package errors

import "fmt"

// InvalidCharacter is created when an invalid character appears.
type InvalidCharacter struct {
	Character string
}

// Error implements the error interface.
func (err *InvalidCharacter) Error() string {
	return fmt.Sprintf("Invalid character '%s'", err.Character)
}