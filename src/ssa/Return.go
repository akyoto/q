package ssa

import (
	"fmt"
	"strings"

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
	if len(v.Arguments) == 0 {
		return "return"
	}

	args := make([]string, 0, len(v.Arguments))

	for _, arg := range v.Arguments {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("return %s", strings.Join(args, ", "))
}

func (v *Return) Type() types.Type {
	return types.Void
}