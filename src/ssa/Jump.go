package ssa

import (
	"fmt"
)

// Jump will transfer program flow to the block it's pointing at.
type Jump struct {
	Void
	To *Block
}

func (a *Jump) Equals(v Value) bool {
	b, sameType := v.(*Jump)

	if !sameType {
		return false
	}

	return a.To == b.To
}

func (v *Jump) Inputs() []Value {
	return nil
}

func (v *Jump) Replace(Value, Value) {}

func (v *Jump) String() string {
	return fmt.Sprintf("jump(%s)", v.To.Label)
}