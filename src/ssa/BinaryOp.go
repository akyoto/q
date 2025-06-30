package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

type BinaryOp struct {
	Left  Value
	Right Value
	Op    token.Kind
	Liveness
	Source
}

func (v *BinaryOp) Dependencies() []Value {
	return []Value{v.Left, v.Right}
}

func (a *BinaryOp) Equals(v Value) bool {
	b, sameType := v.(*BinaryOp)

	if !sameType {
		return false
	}

	if !a.Left.Equals(b.Left) {
		return false
	}

	if !a.Right.Equals(b.Right) {
		return false
	}

	return true
}

func (v *BinaryOp) IsConst() bool {
	return true
}

func (v *BinaryOp) String() string {
	return fmt.Sprintf("%s %s %s", v.Left, expression.Operators[v.Op].Symbol, v.Right)
}

func (v *BinaryOp) Type() types.Type {
	return v.Left.Type()
}