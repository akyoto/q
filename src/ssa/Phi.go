package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Phi is the merging point of multiple values for the same name.
type Phi struct {
	Typ types.Type
	Arguments
	Liveness
}

// Equals returns true if the phi nodes are equal.
func (a *Phi) Equals(v Value) bool {
	b, sameType := v.(*Phi)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsConst returns true because two equal phi nodes produce the same value.
func (p *Phi) IsConst() bool {
	return true
}

// String returns a human-readable representation of the phi node.
func (p *Phi) String() string {
	return fmt.Sprintf("phi(%s)", p.Arguments.String())
}

// Type returns the type of the phi node.
func (p *Phi) Type() types.Type {
	return p.Typ
}