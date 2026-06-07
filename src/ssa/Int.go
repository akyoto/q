package ssa

import (
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

// Int is an integer value.
type Int struct {
	Independent
	Liveness
	Int int
	Source
}

// Equals returns true if the integers are equal.
func (a *Int) Equals(v Value) bool {
	b, sameType := v.(*Int)

	if !sameType {
		return false
	}

	return a.Int == b.Int
}

// IsPure returns true because an integer is always constant.
func (i *Int) IsPure() bool { return true }

// String returns a human-readable representation of the integer.
func (i *Int) String() string { return strconv.Itoa(i.Int) }

// Type returns the type of the integer.
func (i *Int) Type() types.Type { return types.AnyInt }