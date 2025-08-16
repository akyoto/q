package ssa

import (
	"fmt"
)

// Store stores a value at a given index relative to the address.
type Store struct {
	Void
	Address Value
	Index   Value
	Value   Value
	Length  uint8
	Source
}

func (a *Store) Equals(v Value) bool {
	b, sameType := v.(*Store)

	if !sameType {
		return false
	}

	return a.Address == b.Address && a.Index == b.Index && a.Value == b.Value && a.Length == b.Length
}

func (v *Store) Inputs() []Value {
	return []Value{v.Address, v.Index, v.Value}
}

func (v *Store) Replace(old Value, new Value) {
	if v.Address == old {
		v.Address = new
	}

	if v.Index == old {
		v.Index = new
	}

	if v.Value == old {
		v.Value = new
	}
}

func (v *Store) String() string {
	return fmt.Sprintf("store(%db, %p + %p, %p)", v.Length, v.Address, v.Index, v.Value)
}