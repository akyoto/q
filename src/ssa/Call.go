package ssa

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type Call struct {
	Id
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

func (v *Call) Debug() string {
	tmp := strings.Builder{}
	tmp.WriteString("%")
	tmp.WriteString(strconv.Itoa(v.Arguments[0].ID()))
	tmp.WriteString("(")
	args := v.Arguments[1:]

	for i, arg := range args {
		tmp.WriteString("%")
		tmp.WriteString(strconv.Itoa(arg.ID()))

		if i != len(args)-1 {
			tmp.WriteString(", ")
		}
	}

	tmp.WriteString(")")
	return tmp.String()
}

func (v *Call) String() string {
	tmp := strings.Builder{}
	tmp.WriteString(v.Arguments[0].String())
	tmp.WriteString("(")
	args := v.Arguments[1:]

	for i, arg := range args {
		tmp.WriteString(arg.String())

		if i != len(args)-1 {
			tmp.WriteString(", ")
		}
	}

	tmp.WriteString(")")
	return tmp.String()
}

func (v *Call) Type() types.Type {
	typ := v.Arguments[0].(*Function).Typ

	if len(typ.Output) == 0 {
		return types.Void
	}

	return typ.Output[0]
}