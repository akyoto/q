package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Bool is a boolean value.
type Bool struct {
	Structure *Struct
	Liveness
	Bool bool
	Source
}

// Equals returns true if the boolean values are equal.
func (a *Bool) Equals(v Value) bool {
	b, sameType := v.(*Bool)

	if !sameType {
		return false
	}

	return a.Bool == b.Bool
}

// Inputs returns nil because a boolean value has no inputs.
func (b *Bool) Inputs() []Value { return nil }

// IsConst returns true because a boolean value is always constant.
func (b *Bool) IsConst() bool { return true }

// Replace does nothing because a boolean value has no inputs.
func (b *Bool) Replace(Value, Value) {}

// String returns a human-readable representation of the boolean value.
func (b *Bool) String() string { return fmt.Sprint(b.Bool) }

// Struct returns the struct that this boolean value is a part of.
func (b *Bool) Struct() *Struct { return b.Structure }

// Type returns the boolean type.
func (b *Bool) Type() types.Type { return types.Bool }