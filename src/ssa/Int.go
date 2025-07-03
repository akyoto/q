package ssa

import (
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

type Int struct {
	Int int
	Id
	Liveness
	Source
}

func (v *Int) Dependencies() []Value {
	return nil
}

func (a *Int) Equals(v Value) bool {
	b, sameType := v.(*Int)

	if !sameType {
		return false
	}

	return a.Int == b.Int
}

func (v *Int) IsConst() bool {
	return true
}

func (v *Int) Debug(expand bool) string {
	return v.String()
}

func (v *Int) String() string {
	return strconv.Itoa(v.Int)
}

func (v *Int) Type() types.Type {
	return types.AnyInt
}