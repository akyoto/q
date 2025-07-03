package ssa

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type Syscall struct {
	Id
	Arguments
	Liveness
	Source
}

func (a *Syscall) Equals(v Value) bool {
	b, sameType := v.(*Syscall)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Syscall) IsConst() bool {
	return false
}

func (v *Syscall) Debug(expand bool) string {
	tmp := strings.Builder{}
	tmp.WriteString("syscall(")

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

	tmp.WriteString(")")
	return tmp.String()
}

func (v *Syscall) String() string {
	return v.Debug(true)
}

func (v *Syscall) Type() types.Type {
	return types.Any
}