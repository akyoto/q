package ssa

import (
	"fmt"
)

// Assert makes sure that a condition is true at the time of execution.
type Assert struct {
	Void
	Condition Value
}

// Equals returns true if the assert instructions are equal.
func (a *Assert) Equals(v Value) bool {
	b, sameType := v.(*Assert)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition)
}

// Inputs returns the condition.
func (v *Assert) Inputs() []Value {
	return []Value{v.Condition}
}

// Replace replaces the condition if it matches.
func (v *Assert) Replace(old Value, new Value) {
	if v.Condition == old {
		v.Condition = new
	}
}

// String returns a human-readable representation of the assert.
func (v *Assert) String() string {
	return fmt.Sprintf("assert %s", v.Condition.String())
}