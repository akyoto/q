package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// FromTuple is a value inside of a struct or tuple.
type FromTuple struct {
	Tuple Value
	Liveness
	Index int
	Source
}

// Equals returns true if the tuple accesses are equal.
func (a *FromTuple) Equals(v Value) bool {
	b, sameType := v.(*FromTuple)

	if !sameType {
		return false
	}

	return a.Tuple == b.Tuple && a.Index == b.Index
}

// Inputs returns the tuple.
func (f *FromTuple) Inputs() []Value { return []Value{f.Tuple} }

// IsPure returns true because accessing the same element will always have the same value.
func (f *FromTuple) IsPure() bool { return true }

// Replace replaces the tuple if it matches.
func (f *FromTuple) Replace(old Value, new Value) {
	if f.Tuple == old {
		f.Tuple = new
	}
}

// String returns a human-readable representation of the tuple access.
func (f *FromTuple) String() string { return fmt.Sprintf("field(%p, %d)", f.Tuple, f.Index) }

// Type returns the type of the tuple element.
func (f *FromTuple) Type() types.Type {
	switch typ := types.Unwrap(f.Tuple.Type()).(type) {
	case *types.Struct:
		return typ.Fields[f.Index].Type
	case *types.Tuple:
		return typ.Types[f.Index]
	default:
		panic("not implemented")
	}
}