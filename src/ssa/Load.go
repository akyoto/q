package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Load stores a value at a given index relative to the address.
type Load struct {
	Typ     types.Type
	Address Value
	Index   Value
	Liveness
	Source
}

// Equals returns true if the loads are equal.
func (a *Load) Equals(v Value) bool {
	b, sameType := v.(*Load)

	if !sameType {
		return false
	}

	return a.Address == b.Address && a.Index == b.Index
}

// IsConst returns false because a load is a memory access.
func (l *Load) IsConst() bool {
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
	return fmt.Sprintf("load(%db, %p + %p)", l.Typ.Size(), l.Address, l.Index)
}

// Type returns the type of the loaded value.
func (l *Load) Type() types.Type {
	return l.Typ
}