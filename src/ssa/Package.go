package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

type Package struct {
	Name     string
	IsExtern bool
}

func (v *Package) AddUser(Value)    { panic("package can not be used as a dependency") }
func (v *Package) Inputs() []Value  { return nil }
func (v *Package) IsConst() bool    { return true }
func (v *Package) String() string   { return v.Name }
func (v *Package) Type() types.Type { return nil }
func (v *Package) Users() []Value   { return nil }

func (a *Package) Equals(v Value) bool {
	b, sameType := v.(*Package)

	if !sameType {
		return false
	}

	return a.Name == b.Name
}