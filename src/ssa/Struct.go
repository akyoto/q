package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Struct is a list of values that can be addressed by name.
type Struct struct {
	Typ *types.Struct
	Arguments
	Liveness
	Source
}

func (v *Struct) IsConst() bool    { return true }
func (v *Struct) String() string   { return fmt.Sprintf("%s{%s}", v.Typ.Name(), v.Arguments.String()) }
func (v *Struct) Type() types.Type { return v.Typ }

func (a *Struct) Equals(v Value) bool {
	b, sameType := v.(*Struct)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}