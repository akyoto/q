package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Load stores a value at a given index relative to the address.
type Load struct {
	Memory
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

// Inputs returns the address and index of the load.
func (l *Load) Inputs() []Value {
	return []Value{l.Address, l.Index}
}

// Replace replaces the address or index if it matches.
func (l *Load) Replace(old Value, new Value) {
	if l.Address == old {
		l.Address = new
	}

	if l.Index == old {
		l.Index = new
	}
}

// String returns a human-readable representation of the load.
func (l *Load) String() string {
	if l.Scale {
		return fmt.Sprintf("load(%db, %p + %p * %d)", l.Typ.Size(), l.Address, l.Index, l.Typ.Size())
	}

	return fmt.Sprintf("load(%db, %p + %p)", l.Typ.Size(), l.Address, l.Index)
}

// Type returns the type of the loaded value.
func (l *Load) Type() types.Type {
	return l.Typ
}