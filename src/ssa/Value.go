package ssa

import "git.urbach.dev/cli/q/src/types"

type Value interface {
	// AddUser adds a new user of this value.
	AddUser(Value)

	// Equals check for equality which is useful for common subexpression elimination.
	Equals(Value) bool

	// Inputs returns all values that are needed for this value to be calculated.
	Inputs() []Value

	// IsConst returns true if the calculation of the value has no side effects.
	IsConst() bool

	// RemoveUser removes an existing user of this value.
	RemoveUser(Value)

	// Replace replaces the uses of a value with another value.
	Replace(Value, Value)

	// String returns a human-readable form of the value.
	String() string

	// Type returns the data type of the value.
	Type() types.Type

	// Users returns all values that reference this value as an input.
	Users() []Value
}