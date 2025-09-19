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
func (s *Memory) Inputs() []Value {
	return []Value{s.Address, s.Index}
}

// IsPure returns false because the inputs are not consumed.
func (l *Memory) IsPure() bool {
	return false
}

// Replace replaces the address, index, or value if it matches.
func (s *Memory) Replace(old Value, new Value) {
	if s.Address == old {
		s.Address = new
	}

	if s.Index == old {
		s.Index = new
	}
}

// String returns a human-readable representation of the memory.
func (s *Memory) String() string {
	if s.Scale {
		return fmt.Sprintf("memory(%db, %p + %p * %d)", s.Typ.Size(), s.Address, s.Index, s.Typ.Size())
	}

	return fmt.Sprintf("memory(%db, %p + %p)", s.Typ.Size(), s.Address, s.Index)
}

// Type returns the type of the loaded value.
func (s *Memory) Type() types.Type {
	return s.Typ
}