package errors

import "fmt"

// MissingCharacter represents an error where a required character is missing.
type MissingCharacter struct {
	Character string
}

func (err *MissingCharacter) Error() string {
	switch err.Character {
	case "(", "{", "[":
		return fmt.Sprintf("Missing opening bracket: '%s'", err.Character)

	case ")", "}", "]":
		return fmt.Sprintf("Missing closing bracket: '%s'", err.Character)

	default:
		return fmt.Sprintf("Missing character: '%s'", err.Character)
	}
}
