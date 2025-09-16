package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	AlwaysFalse               = errors.String("Condition is always false")
	AlwaysTrue                = errors.String("Condition is always true")
	InvalidCallExpression     = errors.String("Invalid call expression")
	InvalidCondition          = errors.String("Invalid condition")
	InvalidExpression         = errors.String("Invalid expression")
	InvalidFieldInit          = errors.String("Invalid field initialization (expected 'field: value')")
	InvalidLoopHeader         = errors.String("Invalid loop header")
	InvalidNumber             = errors.String("Invalid number")
	InvalidRune               = errors.String("Invalid rune")
	InvalidStructOperation    = errors.String("Invalid operation on structs")
	MissingCommaBetweenFields = errors.String("Missing ',' between struct fields")
	MissingOperand            = errors.String("Missing operand")
	UnnecessaryCast           = errors.String("Unnecessary type cast")
)

// DefinitionCountMismatch error is created when the number of provided definitions doesn't match the return type.
type DefinitionCountMismatch struct {
	Function      string
	Count         int
	ExpectedCount int
}

func (err *DefinitionCountMismatch) Error() string {
	if err.Count > err.ExpectedCount {
		if err.ExpectedCount == 0 {
			return fmt.Sprintf("'%s' does not have a return value", err.Function)
		}

		return fmt.Sprintf("Too many variables for the return value of '%s'", err.Function)
	}

	return fmt.Sprintf("Not enough variables for the return value of '%s'", err.Function)
}

// ErrorNotChecked is created when a variable is accessed without checking its error value.
type ErrorNotChecked struct {
	Identifier string
}

func (err *ErrorNotChecked) Error() string {
	return fmt.Sprintf("Error must be checked before accessing '%s'", err.Identifier)
}

// NoMatchingFunction is created when a function is not defined for the given type.
type NoMatchingFunction struct {
	Function string
}

func (err *NoMatchingFunction) Error() string {
	return fmt.Sprintf("No matching function for call to '%s'", err.Function)
}

// NotDataStruct is created when accessing field of a non-struct type.
type NotDataStruct struct {
	TypeName string
}

func (err *NotDataStruct) Error() string {
	return fmt.Sprintf("Type '%s' is not a data structure", err.TypeName)
}

// NotImplemented represents an error where the implementation is currently missing.
type NotImplemented struct {
	Subject string
}

func (err *NotImplemented) Error() string {
	return fmt.Sprintf("Not implemented: %s", err.Subject)
}

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

// ResourceNotConsumed error is created when a resource has not been consumed in an exit block.
type ResourceNotConsumed struct {
	TypeName string
}

func (err *ResourceNotConsumed) Error() string {
	return fmt.Sprintf("Resource of type '%s' not consumed", err.TypeName)
}

// ResourcePartiallyConsumed error is created when a resource has only partially been consumed in an exit block.
type ResourcePartiallyConsumed struct {
	TypeName string
}

func (err *ResourcePartiallyConsumed) Error() string {
	return fmt.Sprintf("Resource of type '%s' not consumed in all branches", err.TypeName)
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

	return fmt.Sprintf("Expected %s '%s' (encountered '%s')", subject, err.Expected, err.Encountered)
}

// TypeNotIndexable represents an error where a type does not allow indexing.
type TypeNotIndexable struct {
	TypeName string
}

func (err *TypeNotIndexable) Error() string {
	return fmt.Sprintf("Value of type '%s' does not support indexing", err.TypeName)
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

// PartiallyUnknownIdentifier represents identifiers that were only defined in one branch.
type PartiallyUnknownIdentifier struct {
	Name string
}

func (err *PartiallyUnknownIdentifier) Error() string {
	return fmt.Sprintf("Identifier '%s' is not defined in every branch", err.Name)
}

// UndefinedStructField is created when an undefined struct field is accessed.
type UndefinedStructField struct {
	Identifier string
	FieldName  string
}

func (err *UndefinedStructField) Error() string {
	return fmt.Sprintf("Struct field '%s' of '%s' has an undefined value", err.FieldName, err.Identifier)
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

// UnknownType error is created when a type could not be found.
type UnknownType struct {
	Name string
}

func (err *UnknownType) Error() string {
	return fmt.Sprintf("Unknown type '%s'", err.Name)
}

// UnusedValue error is created when a value is never used.
type UnusedValue struct {
	Value string
}

func (err *UnusedValue) Error() string {
	return fmt.Sprintf("Unused value '%s'", err.Value)
}

// VariableAlreadyExists is used when existing variables are used for new variable declarations.
type VariableAlreadyExists struct {
	Name string
}

func (err *VariableAlreadyExists) Error() string {
	return fmt.Sprintf("Variable '%s' already exists", err.Name)
}

// WriteToImmutable error is created when a memory store to an immutable type has been attempted.
type WriteToImmutable struct {
	Name     string
	TypeName string
}

func (err *WriteToImmutable) Error() string {
	return fmt.Sprintf("'%s' of type '%s' is immutable", err.Name, err.TypeName)
}