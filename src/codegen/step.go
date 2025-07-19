package codegen

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

type step struct {
	Value    ssa.Value
	Live     []*step
	Hints    []cpu.Register
	Index    int
	Register cpu.Register
}

func (s *step) Hint(reg cpu.Register) {
	if len(s.Hints) == 0 {
		s.Register = reg
	}

	s.Hints = append(s.Hints, reg)
}

func (s *step) String() string {
	return s.Value.String()
}