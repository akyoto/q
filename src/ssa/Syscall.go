package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Syscall struct {
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

func (v *Syscall) String() string {
	return fmt.Sprintf("syscall(%v)", v.Arguments)
}

func (v *Syscall) Type() types.Type {
	return types.Any
}