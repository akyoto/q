package ssa

import "git.urbach.dev/cli/q/src/types"

type Function struct {
	UniqueName string
	Typ        *types.Function
	IsExtern   bool
	Liveness
	Source
}

func (v *Function) Inputs() []Value  { return nil }
func (v *Function) IsConst() bool    { return true }
func (v *Function) String() string   { return v.UniqueName }
func (v *Function) Type() types.Type { return v.Typ }

func (a *Function) Equals(v Value) bool {
	b, sameType := v.(*Function)

	if !sameType {
		return false
	}

	return a.UniqueName == b.UniqueName
}