package ssa

import (
	"fmt"

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

	if a.Left.Equals(b.Left) && a.Right.Equals(b.Right) {
		return true
	}

	return a.IsCommutative() && a.Left.Equals(b.Right) && a.Right.Equals(b.Left)
}

// Inputs returns the left and right operands.
func (op *BinaryOp) Inputs() []Value {
	return []Value{op.Left, op.Right}
}

// IsAssociative returns true if rearranging the parentheses does not change the result.
func (op *BinaryOp) IsAssociative() bool {
	switch op.Op {
	case token.Add, token.Mul, token.And, token.Or, token.Xor, token.LogicalAnd, token.LogicalOr:
		return true
	default:
		return false
	}
}

// IsCommutative returns true if changing the order of the operands does not change the result.
func (op *BinaryOp) IsCommutative() bool {
	switch op.Op {
	case token.Add, token.Mul, token.And, token.Or, token.Xor, token.LogicalAnd, token.LogicalOr, token.Equal, token.NotEqual:
		return true
	default:
		return false
	}
}

// IsPure returns true if both operands are constant.
func (op *BinaryOp) IsPure() bool {
	return op.Left.IsPure() && op.Right.IsPure()
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
	return fmt.Sprintf("%p %s %p", op.Left, op.Op, op.Right)
}

// Type returns the type of the result of the binary operation.
func (op *BinaryOp) Type() types.Type {
	if op.Op.IsComparison() {
		return types.Bool
	}

	return op.Left.Type()
}