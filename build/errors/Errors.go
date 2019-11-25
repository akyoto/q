package errors

import (
	"fmt"
)

var (
	NotImplemented              = &Base{"Not implemented", false}
	InvalidInstruction          = &Base{"Invalid instruction", false}
	MissingParameter            = &Base{"Missing parameter", false}
	MissingFunctionName         = &Base{"Expected function name before '('", false}
	ExpectedVariable            = &Base{"Expected variable on the left side of the assignment", false}
	InvalidFunctionName         = &Base{"A function can not be named 'func' or 'fn'", false}
	ParameterOpeningBracket     = &Base{"Missing opening bracket '(' after the function name", false}
	TopLevel                    = &Base{"Only function definitions are allowed at the top level", false}
	MissingRange                = &Base{"Missing range expression in for loop", false}
	MissingRangeLimit           = &Base{"Missing upper limit in range expression", true}
	MissingAssignmentOperator   = &Base{"Missing assignment operator", false}
	MissingAssignmentExpression = &Base{"Missing assignment expression", false}
)

// NotANumber represents number conversion errors.
type NotANumber struct {
	Expression string
}

func (err *NotANumber) Error() string {
	return fmt.Sprintf("Not a number: %s", err.Expression)
}

// UnknownVariable represents unknown variables.
type UnknownVariable struct {
	VariableName string
}

func (err *UnknownVariable) Error() string {
	return fmt.Sprintf("Unknown variable: '%s'", err.VariableName)
}

// ParameterCount represents an error where the parameter count is different from the expected number of parameters.
type ParameterCount struct {
	FunctionName  string
	CountGiven    int
	CountRequired int
}

func (err *ParameterCount) Error() string {
	if err.CountGiven < err.CountRequired {
		return fmt.Sprintf("Too few arguments in '%s' call", err.FunctionName)
	}

	if err.CountGiven > err.CountRequired {
		return fmt.Sprintf("Too many arguments in '%s' call", err.FunctionName)
	}

	return ""
}

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
