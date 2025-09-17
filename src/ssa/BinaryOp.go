package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// BinaryOp is an operation with two operands.
type BinaryOp struct {
	Left      Value
	Right     Value
	Structure *Struct
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

// Struct returns the structure this operation belongs to.
func (op *BinaryOp) Struct() *Struct {
	return op.Structure
}

// Type returns the type of the result of the binary operation.
func (op *BinaryOp) Type() types.Type {
	if op.Op.IsComparison() {
		return types.Bool
	}

	return op.Left.Type()
}