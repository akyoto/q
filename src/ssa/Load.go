package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Load stores a value at a given index relative to the address.
type Load struct {
	Memory *Memory
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

// IsPure returns true because loads have no side effects.
// Dead loads can be removed and two equal loads only deduplicate if
// nothing between them modifies the memory.
func (l *Load) IsPure() bool {
	return true
}

// Inputs returns the memory address of the load.
func (l *Load) Inputs() []Value {
	return []Value{l.Memory.Address, l.Memory.Index}
}

// Replace replaces the address or index if it matches.
func (l *Load) Replace(old Value, new Value) {
	l.Memory.Replace(old, new)
}

// String returns a human-readable representation of the load.
func (l *Load) String() string {
	return fmt.Sprintf("load(%s)", l.Memory)
}

// Type returns the type of the loaded value.
func (l *Load) Type() types.Type {
	return l.Memory.Typ
}