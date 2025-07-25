package ssa

import "git.urbach.dev/cli/q/src/types"

type Value interface {
	AddUser(Value)

	// Equals returns true if it's equal to the given value.
	Equals(Value) bool

	// Inputs returns all values that are needed for this value to be calculated.
	Inputs() []Value

	// IsConst returns true if the calculation of the value has no side effects.
	IsConst() bool

	// Replace replaces the uses of a value with another value.
	Replace(Value, Value)

	// Strings returns a human-readable form of the value.
	String() string

	// Type returns the data type.
	Type() types.Type

	// Users returns all values that reference this value as an input.
	Users() []Value
}