package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// Global is a memory location identified by a label.
type Global struct {
	Label string
	Typ   types.Type
	Liveness
	Source
}

// Equals returns true if the data labels are equal.
func (a *Global) Equals(v Value) bool {
	b, sameType := v.(*Global)

	if !sameType {
		return false
	}

	return a.Label == b.Label
}

// Inputs returns nil because a data label has no inputs.
func (d *Global) Inputs() []Value { return nil }

// IsPure returns true because a data label is always constant.
func (d *Global) IsPure() bool { return true }

// Replace does nothing because a data label has no inputs.
func (d *Global) Replace(Value, Value) {}

// String returns a human-readable representation of the data label.
func (d *Global) String() string { return d.Label }

// Type returns the type of the memory it references.
func (d *Global) Type() types.Type { return d.Typ }