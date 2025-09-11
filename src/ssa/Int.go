package ssa

import (
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

// Int is an integer value.
type Int struct {
	Structure *Struct
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

// Inputs returns nil because an integer has no inputs.
func (i *Int) Inputs() []Value { return nil }

// IsPure returns true because an integer is always constant.
func (i *Int) IsPure() bool { return true }

// Replace does nothing because an integer has no inputs.
func (i *Int) Replace(Value, Value) {}

// String returns a human-readable representation of the integer.
func (i *Int) String() string { return strconv.Itoa(i.Int) }

// Struct returns the struct that this integer is a part of.
func (i *Int) Struct() *Struct { return i.Structure }

// Type returns the type of the integer.
func (i *Int) Type() types.Type { return types.AnyInt }