package ssa

import (
	"fmt"
)

// Store stores a value at a given index relative to the address.
type Store struct {
	Void
	Value Value
	Memory
	Source
}

// Equals returns true if the stores are equal.
func (a *Store) Equals(v Value) bool {
	b, sameType := v.(*Store)

	if !sameType {
		return false
	}

	return a.Memory == b.Memory
}

// Inputs returns the address, index, and value of the store.
func (s *Store) Inputs() []Value {
	return []Value{s.Address, s.Index, s.Value}
}

// Replace replaces the address, index, or value if it matches.
func (s *Store) Replace(old Value, new Value) {
	if s.Address == old {
		s.Address = new
	}

	if s.Index == old {
		s.Index = new
	}

	if s.Value == old {
		s.Value = new
	}
}

// String returns a human-readable representation of the store.
func (s *Store) String() string {
	if s.Scale {
		return fmt.Sprintf("store(%db, %p + %p * %d, %p)", s.Typ.Size(), s.Address, s.Index, s.Typ.Size(), s.Value)
	}

	return fmt.Sprintf("store(%db, %p + %p, %p)", s.Typ.Size(), s.Address, s.Index, s.Value)
}