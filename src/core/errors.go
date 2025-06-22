package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	InvalidExpression = errors.String("Invalid expression")
	InvalidNumber     = errors.String("Invalid number")
	InvalidRune       = errors.String("Invalid rune")
)

// UnknownIdentifier represents unknown identifiers.
type UnknownIdentifier struct {
	Name        string
	CorrectName string
}

func (err *UnknownIdentifier) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown identifier '%s', did you mean '%s'?", err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Unknown identifier '%s'", err.Name)
}