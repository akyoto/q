package ssa

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

type Struct struct {
	Typ *types.Struct
	Id
	Arguments
	Liveness
	Source
}

func (a *Struct) Equals(v Value) bool {
	b, sameType := v.(*Struct)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

func (v *Struct) IsConst() bool {
	return true
}

func (v *Struct) Debug() string {
	tmp := strings.Builder{}
	tmp.WriteString(v.Typ.Name())
	tmp.WriteString("{")

	for i, arg := range v.Arguments {
		tmp.WriteString("%")
		tmp.WriteString(strconv.Itoa(arg.ID()))

		if i != len(v.Arguments)-1 {
			tmp.WriteString(", ")
		}
	}

	tmp.WriteString("}")
	return tmp.String()
}

func (v *Struct) String() string {
	tmp := strings.Builder{}
	tmp.WriteString(v.Typ.Name())
	tmp.WriteString("{")

	for i, arg := range v.Arguments {
		tmp.WriteString(arg.String())

		if i != len(v.Arguments)-1 {
			tmp.WriteString(", ")
		}
	}

	tmp.WriteString("}")
	return tmp.String()
}

func (v *Struct) Type() types.Type {
	return v.Typ
}