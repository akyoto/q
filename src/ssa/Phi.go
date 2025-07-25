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

func (v *Phi) IsConst() bool {
	return true
}

func (a *Phi) Equals(v Value) bool {
	b, sameType := v.(*Phi)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Phi) String() string {
	return fmt.Sprintf("phi(%s)", v.Arguments.String())
}

func (v *Phi) Type() types.Type {
	return v.Typ
}