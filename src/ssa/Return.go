package ssa

import (
	"fmt"
)

// Return transfers program flow back to the caller.
type Return struct {
	Void
	Arguments
}

// Equals returns true if the returns are equal.
func (a *Return) Equals(v Value) bool {
	b, sameType := v.(*Return)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// String returns a human-readable representation of the return.
func (r *Return) String() string {
	if len(r.Arguments) == 0 {
		return "return"
	}

	return fmt.Sprintf("return %s", r.Arguments.String())
}