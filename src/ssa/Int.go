package ssa

import (
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

type Int struct {
	Structure *Struct
	Liveness
	Int int
	Source
}

func (v *Int) Inputs() []Value      { return nil }
func (v *Int) IsConst() bool        { return true }
func (v *Int) Replace(Value, Value) {}
func (v *Int) String() string       { return strconv.Itoa(v.Int) }
func (v *Int) Struct() *Struct      { return v.Structure }
func (v *Int) Type() types.Type     { return types.AnyInt }

func (a *Int) Equals(v Value) bool {
	b, sameType := v.(*Int)

	if !sameType {
		return false
	}

	return a.Int == b.Int
}