package ssa

import (
	"fmt"
	"strings"

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
	args := make([]string, 0, len(v.Arguments))

	for _, arg := range v.Arguments {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("syscall(%s)", strings.Join(args, ", "))
}

func (v *Syscall) Type() types.Type {
	return types.Any
}