package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Memory represents a memory address.
type Memory struct {
	Typ     types.Type
	Address Value
	Index   Value
	Scale   bool
	Liveness
	Source
}

// Equals returns true if the memorys are equal.
func (a *Memory) Equals(v Value) bool {
	b, sameType := v.(*Memory)

	if !sameType {
		return false
	}

	return a.Address == b.Address && a.Index == b.Index && a.Scale == b.Scale && a.Typ == b.Typ
}

// Inputs returns the address and index of the memory.
func (m *Memory) Inputs() []Value {
	return []Value{m.Address, m.Index}
}

// IsPure returns false because the inputs are not consumed.
func (m *Memory) IsPure() bool {
	return false
}

// Replace replaces the address, index, or value if it matches.
func (m *Memory) Replace(old Value, new Value) {
	if m.Address == old {
		m.Address = new
	}

	if m.Index == old {
		m.Index = new
	}
}

// String returns a human-readable representation of the memory.
func (m *Memory) String() string {
	if m.Scale {
		return fmt.Sprintf("memory(%db, %p + %p * %d)", m.Typ.Size(), m.Address, m.Index, m.Typ.Size())
	}

	return fmt.Sprintf("memory(%db, %p + %p)", m.Typ.Size(), m.Address, m.Index)
}

// Type returns the type of the loaded value.
func (m *Memory) Type() types.Type {
	return m.Typ
}