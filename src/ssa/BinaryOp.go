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

func (v *BinaryOp) Inputs() []Value { return []Value{v.Left, v.Right} }
func (v *BinaryOp) IsConst() bool   { return v.Left.IsConst() && v.Right.IsConst() }

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

func (v *BinaryOp) Replace(old Value, new Value) {
	if v.Left == old {
		v.Left = new
	}

	if v.Right == old {
		v.Right = new
	}
}

func (v *BinaryOp) String() string {
	return fmt.Sprintf("%s %s %s", v.Left, expression.Operators[v.Op].Symbol, v.Right)
}

func (v *BinaryOp) Type() types.Type {
	if v.Op.IsComparison() {
		return types.Bool
	}

	return v.Left.Type()
}