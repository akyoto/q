package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Call struct {
	Func *Function
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
	return fmt.Sprintf("%s(%s)", v.Func.String(), v.Arguments.String())
}

func (v *Call) Type() types.Type {
	if len(v.Func.Typ.Output) == 0 {
		return types.Void
	}

	return v.Func.Typ.Output[0]
}