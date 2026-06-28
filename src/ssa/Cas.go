package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Cas performs an atomic compare and swap.
type Cas struct {
	Arguments
	Liveness
	Source
}

// Equals returns true if the CAS arguments are equal.
func (a *Cas) Equals(v Value) bool {
	b, sameType := v.(*Cas)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsPure returns false because a CAS operation has side effects.
func (c *Cas) IsPure() bool { return false }

// String returns a human-readable representation of the CAS operation.
func (c *Cas) String() string { return fmt.Sprintf("cas(%s)", c.Arguments.String()) }

// Type returns the type of the CAS operation.
func (c *Cas) Type() types.Type { return types.UInt32 }