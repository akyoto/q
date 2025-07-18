package ssa

import "git.urbach.dev/cli/q/src/types"

type Value interface {
	// Type returns the data type.
	Type() types.Type

	// Users returns all values that reference this value as an input.
	Users() []Value

	// Inputs returns all values that are needed for this value to be calculated.
	Inputs() []Value

	// Equals returns true if it's equal to the given value.
	Equals(Value) bool

	// IsConst returns true if the calculation of the value has no side effects.
	IsConst() bool

	// Strings returns a human-readable form of the value.
	String() string
}