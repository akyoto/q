package ssa

import "git.urbach.dev/cli/q/src/cpu"

type Register struct {
	Register cpu.Register
	Liveness
	HasToken
}

func (v *Register) Dependencies() []Value {
	return nil
}

func (a *Register) Equals(v Value) bool {
	b, sameType := v.(*Register)

	if !sameType {
		return false
	}

	return a.Register == b.Register
}

func (v *Register) IsConst() bool {
	return true
}

func (v *Register) String() string {
	return v.Register.String()
}