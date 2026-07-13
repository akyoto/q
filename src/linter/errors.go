package linter

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
)

var (
	AlwaysFalse    = errors.String("Condition is always false")
	AlwaysTrue     = errors.String("Condition is always true")
	DivisionByZero = errors.String("Division by zero")
)

// IdenticalExpressions error is created when a binary operation uses identical expressions
// with an operator that makes the entire operation pointless.
type IdenticalExpressions struct {
	Operator string
}

func (err *IdenticalExpressions) Error() string {
	return fmt.Sprintf("Identical expressions to the left and right of the '%s' operator", err.Operator)
}

// MixedSignedUnsigned represents an error where unsigned and signed types are used.
type MixedSignedUnsigned struct {
	Signed   string
	Unsigned string
}

func (err *MixedSignedUnsigned) Error() string {
	return fmt.Sprintf("Mixed signed '%s' and unsigned '%s'", err.Signed, err.Unsigned)
}

// Simplify error is created when a value can be replaced with another one.
type Simplify struct {
	To string
}

func (err *Simplify) Error() string {
	return fmt.Sprintf("Can be simplified to '%s'", err.To)
}