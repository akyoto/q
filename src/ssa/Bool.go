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

func (v *Bool) Inputs() []Value      { return nil }
func (v *Bool) IsConst() bool        { return true }
func (v *Bool) Replace(Value, Value) {}
func (v *Bool) String() string       { return fmt.Sprint(v.Bool) }
func (v *Bool) Struct() *Struct      { return v.Structure }
func (v *Bool) Type() types.Type     { return types.Bool }

func (a *Bool) Equals(v Value) bool {
	b, sameType := v.(*Bool)

	if !sameType {
		return false
	}

	return a.Bool == b.Bool
}