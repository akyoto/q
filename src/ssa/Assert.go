package ssa

import (
	"fmt"
)

// Assert makes sure that a condition is true at the time of execution.
type Assert struct {
	Condition Value
	Source
	Void
}

func (a *Assert) Equals(v Value) bool {
	b, sameType := v.(*Assert)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition)
}

func (v *Assert) Inputs() []Value {
	return []Value{v.Condition}
}

func (v *Assert) Replace(old Value, new Value) {
	if v.Condition == old {
		v.Condition = new
	}
}

func (v *Assert) String() string {
	return fmt.Sprintf("assert %s", v.Condition.String())
}