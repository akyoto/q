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
	InvalidTopLevel            = errors.String("Invalid top level instruction")
	MissingAssign              = errors.String("Missing '='")
	MissingAssignOrBlock       = errors.String("Missing '=' or '{'")
	MissingBlockStart          = errors.String("Missing '{'")
	MissingBlockEnd            = errors.String("Missing '}'")
	MissingExpression          = errors.String("Missing expression")
	MissingGroupStart          = errors.String("Missing '('")
	MissingGroupEnd            = errors.String("Missing ')'")
	MissingParameter           = errors.String("Missing parameter")
	MissingParameterType       = errors.String("Missing parameter type")
	NoInputFiles               = errors.String("No input files")
)

// InvalidCharacter is created when an invalid character appears.
type InvalidCharacter struct {
	Character string
}

func (err *InvalidCharacter) Error() string {
	return fmt.Sprintf("Invalid character '%s'", err.Character)
}

// IsNotDirectory error is created when a path is not a directory.
type IsNotDirectory struct {
	Path string
}

func (err *IsNotDirectory) Error() string {
	return fmt.Sprintf("'%s' is not a directory", err.Path)
}

// UnexpectedIdentifier error is created when 'func' or 'fn' is found at the top level.
type UnexpectedIdentifier struct {
	Keyword string
	Name    string
}

func (err *UnexpectedIdentifier) Error() string {
	return fmt.Sprintf("Q does not use the '%s' keyword: write %s() {} instead", err.Keyword, err.Name)
}

// UnknownImport error is created when a package import failed.
type UnknownImport struct {
	Package string
}

func (err *UnknownImport) Error() string {
	return fmt.Sprintf("Package '%s' does not exist", err.Package)
}