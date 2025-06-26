package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Return struct {
	Arguments
	Source
}

func (a *Return) AddUse(user Value) { panic("return is not a value") }
func (a *Return) Alive() int        { return 0 }

func (a *Return) Equals(v Value) bool {
	b, sameType := v.(*Return)

	if !sameType {
		return false
	}

	if len(a.Arguments) != len(b.Arguments) {
		return false
	}

	for i := range a.Arguments {
		if !a.Arguments[i].Equals(b.Arguments[i]) {
			return false
		}
	}

	return true
}

func (v *Return) IsConst() bool {
	return false
}

func (v *Return) String() string {
	return fmt.Sprintf("return %v", v.Arguments)
}

func (v *Return) Type() types.Type {
	return types.Void
}