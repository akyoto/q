package ssa

import (
	"fmt"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type Call struct {
	Arguments
	Liveness
	Source
}

func (a *Call) Equals(v Value) bool {
	b, sameType := v.(*Call)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Call) IsConst() bool {
	return false
}

func (v *Call) String() string {
	args := make([]string, 0, len(v.Arguments)-1)

	for _, arg := range v.Arguments[1:] {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%s(%s)", v.Arguments[0].String(), strings.Join(args, ", "))
}

func (v *Call) Type() types.Type {
	typ := v.Arguments[0].(*Function).Typ

	if len(typ.Output) == 0 {
		return types.Void
	}

	return typ.Output[0]
}