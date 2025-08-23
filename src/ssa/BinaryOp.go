package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// BinaryOp is an operation with two operands.
type BinaryOp struct {
	Left  Value
	Right Value
	Liveness
	Source
	Op token.Kind
}

// Equals returns true if the binary operations are equal.
func (a *BinaryOp) Equals(v Value) bool {
	b, sameType := v.(*BinaryOp)

	if !sameType {
		return false
	}

	if a.Op != b.Op {
		return false
	}

	return a.Left.Equals(b.Left) && a.Right.Equals(b.Right)
}

// Inputs returns the left and right operands.
func (op *BinaryOp) Inputs() []Value {
	return []Value{op.Left, op.Right}
}

// IsConst returns true if both operands are constant.
func (op *BinaryOp) IsConst() bool {
	return op.Left.IsConst() && op.Right.IsConst()
}

// Replace replaces the left or right operand if it matches.
func (op *BinaryOp) Replace(old Value, new Value) {
	if op.Left == old {
		op.Left = new
	}

	if op.Right == old {
		op.Right = new
	}
}

// String returns a human-readable representation of the binary operation.
func (op *BinaryOp) String() string {
	return fmt.Sprintf("%p %s %p", op.Left, expression.Operators[op.Op].Symbol, op.Right)
}

// Type returns the type of the result of the binary operation.
func (op *BinaryOp) Type() types.Type {
	if op.Op.IsComparison() {
		return types.Bool
	}

	return op.Left.Type()
}