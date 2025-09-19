package ssa

import (
	"fmt"
)

// Store stores a value at a given index relative to the address.
type Store struct {
	Void
	Value  Value
	Memory Value
	Source
}

// Equals returns true if the stores are equal.
func (a *Store) Equals(v Value) bool {
	b, sameType := v.(*Store)

	if !sameType {
		return false
	}

	return a.Memory == b.Memory && a.Value == b.Value
}

// Inputs returns the memory address and value of the store.
func (s *Store) Inputs() []Value {
	return []Value{s.Memory, s.Value}
}

// Replace replaces the address, index, or value if it matches.
func (s *Store) Replace(old Value, new Value) {
	if s.Memory == old {
		s.Memory = new
	}

	if s.Value == old {
		s.Value = new
	}
}

// String returns a human-readable representation of the store.
func (s *Store) String() string {
	return fmt.Sprintf("store(%p, %p)", s.Memory, s.Value)
}