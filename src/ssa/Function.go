package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Function struct {
	Typ     *types.Function
	Package string
	Name    string
	Liveness
	Source
	IsExtern bool
}

func (v *Function) Inputs() []Value      { return nil }
func (v *Function) IsConst() bool        { return true }
func (v *Function) Replace(Value, Value) {}
func (v *Function) Type() types.Type     { return v.Typ }

func (a *Function) Equals(v Value) bool {
	b, sameType := v.(*Function)

	if !sameType {
		return false
	}

	return a.Package == b.Package && a.Name == b.Name
}

func (v *Function) String() string {
	if v.Package == "" {
		return v.Name
	}

	return fmt.Sprintf("%s.%s", v.Package, v.Name)
}