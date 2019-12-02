package errors

import "fmt"

// InvalidCharacter represents an error where a required character is missing.
type InvalidCharacter struct {
	Character string
}

func (err *InvalidCharacter) Error() string {
	return fmt.Sprintf("Invalid character '%s'", err.Character)
}
