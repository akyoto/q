package ssa

import (
	"fmt"
)

// Branch is a conditional branch.
type Branch struct {
	Void
	Condition Value
	Then      *Block
	Else      *Block
}

// Equals returns true if the branches are equal.
func (a *Branch) Equals(v Value) bool {
	b, sameType := v.(*Branch)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition) && a.Then == b.Then && a.Else == b.Else
}

// Inputs returns the condition.
func (b *Branch) Inputs() []Value {
	return []Value{b.Condition}
}

// Replace replaces the condition if it matches.
func (b *Branch) Replace(old Value, new Value) {
	if b.Condition == old {
		b.Condition = new
	}
}

// String returns a human-readable representation of the branch.
func (b *Branch) String() string {
	return fmt.Sprintf("branch(%p, %s, %s)", b.Condition, b.Then.Label, b.Else.Label)
}