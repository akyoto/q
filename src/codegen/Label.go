package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

type Label struct {
	ssa.Void
	Name string
}

func (v *Label) Inputs() []ssa.Value          { return nil }
func (v *Label) IsConst() bool                { return false }
func (v *Label) String() string               { return v.Name }
func (v *Label) Replace(ssa.Value, ssa.Value) {}
func (v *Label) Type() types.Type             { return types.Void }

func (a *Label) Equals(v ssa.Value) bool {
	b, sameType := v.(*Label)

	if !sameType {
		return false
	}

	return a.Name == b.Name
}