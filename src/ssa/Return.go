package ssa

import (
	"fmt"
)

// Return transfers program flow back to the caller.
type Return struct {
	Void
	Arguments
	Source
}

func (a *Return) Equals(v Value) bool {
	b, sameType := v.(*Return)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Return) String() string {
	if len(v.Arguments) == 0 {
		return "return"
	}

	return fmt.Sprintf("return %s", v.Arguments.String())
}