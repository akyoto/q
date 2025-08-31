package ssa

import (
	"fmt"
)

// Jump will transfer program flow to the block it's pointing at.
type Jump struct {
	Void
	To *Block
}

// Equals returns true if the jumps are equal.
func (a *Jump) Equals(v Value) bool {
	b, sameType := v.(*Jump)

	if !sameType {
		return false
	}

	return a.To == b.To
}

// Inputs returns nil because a jump has no inputs.
func (j *Jump) Inputs() []Value {
	return nil
}

// Replace does nothing because a jump has no inputs.
func (j *Jump) Replace(Value, Value) {}

// String returns a human-readable representation of the jump.
func (j *Jump) String() string {
	return fmt.Sprintf("jump(%s)", CleanLabel(j.To.Label))
}