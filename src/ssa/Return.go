package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Return struct {
	Arguments
	Source
}

func (v *Return) AddUser(Value)    { panic("return can not be used as a dependency") }
func (v *Return) IsConst() bool    { return false }
func (v *Return) Type() types.Type { return types.Void }
func (v *Return) Users() []Value   { return nil }

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

func (v *Return) String() string {
	if len(v.Arguments) == 0 {
		return "return"
	}

	return fmt.Sprintf("return %s", v.Arguments.String())
}