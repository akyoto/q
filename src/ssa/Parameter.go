package ssa

import "fmt"

type Parameter struct {
	Index uint8
	Liveness
	HasToken
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

func (v *Parameter) String() string {
	return fmt.Sprintf("arg[%d]", v.Index)
}