package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/expression"
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

func (v *UnaryOp) Inputs() []Value { return []Value{v.Operand} }
func (v *UnaryOp) IsConst() bool   { return v.Operand.IsConst() }

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

func (v *UnaryOp) Replace(old Value, new Value) {
	if v.Operand == old {
		v.Operand = new
	}
}

func (v *UnaryOp) String() string {
	return fmt.Sprintf("%s(%s)", expression.Operators[v.Op].Symbol, v.Operand)
}

func (v *UnaryOp) Type() types.Type {
	return v.Operand.Type()
}