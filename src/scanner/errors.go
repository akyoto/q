package scanner

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	ExpectedFunctionDefinition = errors.String("Expected function definition")
	ExpectedPackageName        = errors.String("Expected package name")
	InvalidExpression          = errors.String("Invalid expression")
	InvalidFunctionDefinition  = errors.String("Invalid function definition")
	InvalidParameterName       = errors.String("Invalid parameter name")
	MissingBlockStart          = errors.String("Missing '{'")
	MissingBlockEnd            = errors.String("Missing '}'")
	MissingExpression          = errors.String("Missing expression")
	MissingGroupStart          = errors.String("Missing '('")
	MissingGroupEnd            = errors.String("Missing ')'")
	MissingParameter           = errors.String("Missing parameter")
	MissingParameterType       = errors.String("Missing parameter type")
	NoInputFiles               = errors.String("No input files")
)

// UnknownImport error is created when a package import failed.
type UnknownImport struct {
	Package string
}

func (err *UnknownImport) Error() string {
	return fmt.Sprintf("Unknown import '%s'", err.Package)
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