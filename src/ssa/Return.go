package ssa

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type Return struct {
	Id
	Arguments
	Source
	NoLiveness
}

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

func (v *Return) Debug(expand bool) string {
	if len(v.Arguments) == 0 {
		return "return"
	}

	tmp := strings.Builder{}
	tmp.WriteString("return ")

	for i, arg := range v.Arguments {
		if expand {
			tmp.WriteString(arg.String())
		} else {
			tmp.WriteString("%")
			tmp.WriteString(strconv.Itoa(arg.ID()))
		}

		if i != len(v.Arguments)-1 {
			tmp.WriteString(", ")
		}
	}

	return tmp.String()
}

func (v *Return) String() string {
	return v.Debug(true)
}

func (v *Return) Type() types.Type {
	return types.Void
}