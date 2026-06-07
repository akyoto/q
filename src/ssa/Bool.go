package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Bool is a boolean value.
type Bool struct {
	Independent
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

// IsPure returns true because a boolean value is always constant.
func (b *Bool) IsPure() bool { return true }

// String returns a human-readable representation of the boolean value.
func (b *Bool) String() string { return fmt.Sprint(b.Bool) }

// Type returns the boolean type.
func (b *Bool) Type() types.Type { return types.Bool }