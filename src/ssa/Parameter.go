package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Parameter struct {
	Index uint8
	Name  string
	Typ   types.Type
	Id
	Liveness
	Source
}

func (v *Parameter) Dependencies() []Value {
	return nil
}

func (a *Parameter) Equals(v Value) bool {
	b, sameType := v.(*Parameter)

	if !sameType {
		return false
	}

	return a.Index == b.Index
}

func (v *Parameter) IsConst() bool {
	return true
}

func (v *Parameter) Debug(expand bool) string {
	return v.String()
}

func (v *Parameter) String() string {
	return fmt.Sprintf("args[%d]", v.Index)
}

func (v *Parameter) Type() types.Type {
	return v.Typ
}