package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

type BinaryOperation struct {
	Left  Value
	Right Value
	Op    token.Kind
	Liveness
	HasToken
}

func (v *BinaryOperation) Dependencies() []Value {
	return []Value{v.Left, v.Right}
}

func (a *BinaryOperation) Equals(v Value) bool {
	b, sameType := v.(*BinaryOperation)

	if !sameType {
		return false
	}

	if a.Source.Kind != b.Source.Kind {
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

func (v *BinaryOperation) IsConst() bool {
	return true
}

func (v *BinaryOperation) String() string {
	return fmt.Sprintf("%s %s %s", v.Left, expression.Operators[v.Op].Symbol, v.Right)
}