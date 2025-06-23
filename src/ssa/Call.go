package ssa

import "fmt"

type Call struct {
	Arguments
	Liveness
	HasToken
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

func (v *Call) String() string {
	return fmt.Sprintf("call%v", v.Args)
}