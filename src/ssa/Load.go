package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Load stores a value at a given index relative to the address.
type Load struct {
	Memory Value
	Liveness
	Source
}

// Equals returns true if the loads are equal.
func (a *Load) Equals(v Value) bool {
	b, sameType := v.(*Load)

	if !sameType {
		return false
	}

	return a.Memory == b.Memory
}

// IsPure returns false because loads read from memory
// and two equal loads may yield different values.
func (l *Load) IsPure() bool {
	return false
}

// Inputs returns the memory address of the load.
func (l *Load) Inputs() []Value {
	return []Value{l.Memory}
}

// Replace replaces the address or index if it matches.
func (l *Load) Replace(old Value, new Value) {
	if l.Memory == old {
		l.Memory = new
	}
}

// String returns a human-readable representation of the load.
func (l *Load) String() string {
	return fmt.Sprintf("load(%p)", l.Memory)
}

// Type returns the type of the loaded value.
func (l *Load) Type() types.Type {
	return l.Memory.(*Memory).Typ
}