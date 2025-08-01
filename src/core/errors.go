package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	InvalidCondition  = errors.String("Invalid condition")
	InvalidExpression = errors.String("Invalid expression")
	InvalidLoopHeader = errors.String("Invalid loop header")
	InvalidNumber     = errors.String("Invalid number")
	InvalidRune       = errors.String("Invalid rune")
	MissingOperand    = errors.String("Missing operand")
)

// ParameterCountMismatch error is created when the number of provided parameters doesn't match the function signature.
type ParameterCountMismatch struct {
	Function      string
	Count         int
	ExpectedCount int
}

func (err *ParameterCountMismatch) Error() string {
	if err.Count > err.ExpectedCount {
		return fmt.Sprintf("Too many parameters in '%s' function call", err.Function)
	}

	return fmt.Sprintf("Not enough parameters in '%s' function call", err.Function)
}

// ReturnCountMismatch error is created when the number of returned values doesn't match the return type.
type ReturnCountMismatch struct {
	Count         int
	ExpectedCount int
}

func (err *ReturnCountMismatch) Error() string {
	if err.Count > err.ExpectedCount {
		return fmt.Sprintf("Too many return values (expected %d)", err.ExpectedCount)
	}

	return fmt.Sprintf("Not enough return values (expected %d)", err.ExpectedCount)
}

// TypeMismatch represents an error where a type requirement was not met.
type TypeMismatch struct {
	Encountered   string
	Expected      string
	ParameterName string
	IsReturn      bool
}

func (err *TypeMismatch) Error() string {
	subject := "type"

	if err.IsReturn {
		subject = "return type"
	}

	if err.ParameterName != "" {
		return fmt.Sprintf("Expected parameter '%s' of %s '%s' (encountered '%s')", err.ParameterName, subject, err.Expected, err.Encountered)
	}

	return fmt.Sprintf("Expected %s '%s' instead of '%s'", subject, err.Expected, err.Encountered)
}

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

// UnknownStructField represents unknown struct fields.
type UnknownStructField struct {
	StructName       string
	FieldName        string
	CorrectFieldName string
}

func (err *UnknownStructField) Error() string {
	if err.CorrectFieldName != "" {
		return fmt.Sprintf("Unknown struct field '%s' in '%s', did you mean '%s'?", err.FieldName, err.StructName, err.CorrectFieldName)
	}

	return fmt.Sprintf("Unknown struct field '%s' in '%s'", err.FieldName, err.StructName)
}

// UnusedValue error is created when a value is never used.
type UnusedValue struct {
	Value string
}

func (err *UnusedValue) Error() string {
	return fmt.Sprintf("Unused value '%s'", err.Value)
}