package ssa2asm

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

type Step struct {
	Index    int
	Value    ssa.Value
	Live     []*Step
	Hints    []cpu.Register
	Register cpu.Register
}

func (s *Step) Hint(reg cpu.Register) {
	if len(s.Hints) == 0 {
		s.Register = reg
	}

	s.Hints = append(s.Hints, reg)
}

func (s *Step) String() string {
	return s.Value.String()
}