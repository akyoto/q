package build

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented          = errors.New("Not implemented")
	ErrInvalidInstruction      = errors.New("Invalid instruction")
	ErrMissingParameter        = errors.New("Missing parameter")
	ErrMissingFunctionName     = errors.New("Expected function name before '('")
	ErrExpectedVariable        = errors.New("Expected variable on the left side of the assignment")
	ErrInvalidFunctionName     = errors.New("A function can not be named 'func' or 'fn'")
	ErrParameterOpeningBracket = errors.New("Missing opening bracket '(' after the function name")
	ErrTopLevel                = errors.New("Only function definitions are allowed at the top level")
	ErrMissingRange            = errors.New("Missing upper limit in for loop")
)

// ErrNotANumber represents number conversion errors.
type ErrNotANumber struct {
	Expression string
}

func (err *ErrNotANumber) Error() string {
	return fmt.Sprintf("Not a number: %s", err.Expression)
}

// ErrUnknownVariable represents unknown variables.
type ErrUnknownVariable struct {
	VariableName string
}

func (err *ErrUnknownVariable) Error() string {
	return fmt.Sprintf("Unknown variable: '%s'", err.VariableName)
}

// ErrParameterCount represents an error where the parameter count is different from the expected number of parameters.
type ErrParameterCount struct {
	Call FunctionCall
}

func (err *ErrParameterCount) Error() string {
	if len(err.Call.Parameters) < len(err.Call.Function.Parameters) {
		return fmt.Sprintf("Too few arguments in '%s' call", err.Call.Function.Name)
	}

	if len(err.Call.Parameters) > len(err.Call.Function.Parameters) {
		return fmt.Sprintf("Too many arguments in '%s' call", err.Call.Function.Name)
	}

	return ""
}

// ErrMissingCharacter represents an error where a required character is missing.
type ErrMissingCharacter struct {
	Character string
}

func (err *ErrMissingCharacter) Error() string {
	switch err.Character {
	case "(", "{", "[":
		return fmt.Sprintf("Missing opening bracket: '%s'", err.Character)

	case ")", "}", "]":
		return fmt.Sprintf("Missing closing bracket: '%s'", err.Character)

	default:
		return fmt.Sprintf("Missing character: '%s'", err.Character)
	}
}
