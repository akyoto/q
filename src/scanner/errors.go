package scanner

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	ExpectedFunctionDefinition = errors.String("Expected function definition")
	ExpectedPackageName        = errors.String("Expected package name")
	InvalidFunctionDefinition  = errors.String("Invalid function definition")
	MissingBlockStart          = errors.String("Missing '{'")
	MissingBlockEnd            = errors.String("Missing '}'")
	MissingGroupStart          = errors.String("Missing '('")
	MissingGroupEnd            = errors.String("Missing ')'")
	MissingParameter           = errors.String("Missing parameter")
	MissingType                = errors.String("Missing type")
	NoInputFiles               = errors.String("No input files")
)

// CouldNotImport error is created when a package import failed.
type CouldNotImport struct {
	Package string
}

func (err *CouldNotImport) Error() string {
	return fmt.Sprintf("Could not import '%s'", err.Package)
}

// InvalidCharacter is created when an invalid character appears.
type InvalidCharacter struct {
	Character string
}

func (err *InvalidCharacter) Error() string {
	return fmt.Sprintf("Invalid character '%s'", err.Character)
}

// InvalidTopLevel error is created when a top-level instruction is not valid.
type InvalidTopLevel struct {
	Instruction string
}

func (err *InvalidTopLevel) Error() string {
	return fmt.Sprintf("Invalid top level instruction '%s'", err.Instruction)
}

// IsNotDirectory error is created when a path is not a directory.
type IsNotDirectory struct {
	Path string
}

func (err *IsNotDirectory) Error() string {
	return fmt.Sprintf("'%s' is not a directory", err.Path)
}