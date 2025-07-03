package ssa

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type CallExtern struct {
	Id
	Arguments
	Liveness
	Source
}

func (a *CallExtern) Equals(v Value) bool {
	b, sameType := v.(*CallExtern)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *CallExtern) IsConst() bool {
	return false
}

func (v *CallExtern) Debug(expand bool) string {
	tmp := strings.Builder{}

	if expand {
		tmp.WriteString(v.Arguments[0].String())
	} else {
		tmp.WriteString("%")
		tmp.WriteString(strconv.Itoa(v.Arguments[0].ID()))
	}

	tmp.WriteString("(")
	args := v.Arguments[1:]

	for i, arg := range args {
		if expand {
			tmp.WriteString(arg.String())
		} else {
			tmp.WriteString("%")
			tmp.WriteString(strconv.Itoa(arg.ID()))
		}

		if i != len(args)-1 {
			tmp.WriteString(", ")
		}
	}

	tmp.WriteString(")")
	return tmp.String()
}

func (v *CallExtern) String() string {
	return v.Debug(true)
}

func (v *CallExtern) Type() types.Type {
	typ := v.Arguments[0].(*Function).Typ

	if len(typ.Output) == 0 {
		return types.Void
	}

	return typ.Output[0]
}