package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// Data is a memory location identified by a label.
type Data struct {
	Label string
	Typ   types.Type
	Liveness
	Source
}

// Equals returns true if the data labels are equal.
func (a *Data) Equals(v Value) bool {
	b, sameType := v.(*Data)

	if !sameType {
		return false
	}

	return a.Label == b.Label
}

// Inputs returns nil because a data label has no inputs.
func (d *Data) Inputs() []Value { return nil }

// IsPure returns true because a data label is always constant.
func (d *Data) IsPure() bool { return true }

// Replace does nothing because a data label has no inputs.
func (d *Data) Replace(Value, Value) {}

// String returns a human-readable representation of the data label.
func (d *Data) String() string { return d.Label }

// Type returns the type of the memory it references.
func (d *Data) Type() types.Type { return d.Typ }