package ssa

import "git.urbach.dev/cli/q/src/types"

type Value interface {
	// Equals check for equality which is useful for common subexpression elimination.
	Equals(Value) bool

	// Inputs returns all values that are needed for this value to be calculated.
	Inputs() []Value

	// IsPure returns true if the calculation of the value has no side effects.
	IsPure() bool

	// Replace replaces the uses of a value with another value.
	Replace(Value, Value)

	// String returns a human-readable form of the value.
	String() string

	// Type returns the data type of the value.
	Type() types.Type

	// HasUsers is an interface to track where this value is used.
	HasUsers
}