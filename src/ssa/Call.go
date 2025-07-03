package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Call struct {
	Arguments
	Liveness
	Source
}

func (v *Call) IsConst() bool { return false }

func (a *Call) Equals(v Value) bool {
	b, sameType := v.(*Call)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Call) String() string {
	return fmt.Sprintf("%s(%s)", v.Arguments[0].String(), v.Arguments[1:].String())
}

func (v *Call) Type() types.Type {
	typ := v.Arguments[0].(*Function).Typ

	if len(typ.Output) == 0 {
		return types.Void
	}

	return typ.Output[0]
}