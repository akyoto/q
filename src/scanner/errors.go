package scanner

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	expectedPackageName = &errors.String{Message: "Expected package name"}
)

// invalidCharacter is created when an invalid character appears.
type invalidCharacter struct {
	Character string
}

// Error implements the error interface.
func (err *invalidCharacter) Error() string {
	return fmt.Sprintf("Invalid character '%s'", err.Character)
}

// invalidTopLevel error is created when a top-level instruction is not valid.
type invalidTopLevel struct {
	Instruction string
}

// Error generates the string representation.
func (err *invalidTopLevel) Error() string {
	return fmt.Sprintf("Invalid top level instruction '%s'", err.Instruction)
}

// isNotDirectory error is created when a path is not a directory.
type isNotDirectory struct {
	Path string
}

// Error generates the string representation.
func (err *isNotDirectory) Error() string {
	return fmt.Sprintf("Import path '%s' is not a directory", err.Path)
}