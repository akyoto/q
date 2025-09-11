package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// UnaryOp is an operation with a single operand.
type UnaryOp struct {
	Operand Value
	Liveness
	Source
	Op token.Kind
}

// Equals returns true if the unary operations are equal.
func (a *UnaryOp) Equals(v Value) bool {
	b, sameType := v.(*UnaryOp)

	if !sameType {
		return false
	}

	if a.Op != b.Op {
		return false
	}

	return a.Operand.Equals(b.Operand)
}

// Inputs returns the operand.
func (u *UnaryOp) Inputs() []Value { return []Value{u.Operand} }

// IsPure returns true if the operand is constant.
func (u *UnaryOp) IsPure() bool { return u.Operand.IsPure() }

// Replace replaces the operand if it matches.
func (u *UnaryOp) Replace(old Value, new Value) {
	if u.Operand == old {
		u.Operand = new
	}
}

// String returns a human-readable representation of the unary operation.
func (u *UnaryOp) String() string {
	return fmt.Sprintf("%s(%s)", u.Op, u.Operand)
}

// Type returns the type of the operand.
func (u *UnaryOp) Type() types.Type {
	return u.Operand.Type()
}