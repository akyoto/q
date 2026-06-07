package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// Global is a memory location identified by a label.
type Global struct {
	Label string
	Typ   types.Type
	Independent
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

// IsPure returns true because a data label is always constant.
func (d *Global) IsPure() bool { return true }

// String returns a human-readable representation of the data label.
func (d *Global) String() string { return d.Label }

// Type returns the type of the memory it references.
func (d *Global) Type() types.Type { return d.Typ }