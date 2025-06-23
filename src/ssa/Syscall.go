package ssa

import "fmt"

type Syscall struct {
	Arguments
	Liveness
	HasToken
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

func (v *Syscall) String() string {
	return fmt.Sprintf("syscall%v", v.Args)
}