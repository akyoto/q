package codegen

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// Step is created for every single SSA value and holds additional metadata.
type Step struct {
	Value    ssa.Value
	Block    *ssa.Block
	Phi      *Step
	Live     []*Step
	Hints    []cpu.Register
	Index    int
	Register cpu.Register
}

// hint adds a register hint to the step.
func (s *Step) hint(reg cpu.Register) {
	if len(s.Hints) == 0 {
		s.Register = reg
	}

	s.Hints = append(s.Hints, reg)
}

// needsRegister returns true if the value requires a register.
func (s *Step) needsRegister() bool {
	typ := s.Value.Type()

	if typ == types.Void {
		return false
	}

	_, isStruct := typ.(*types.Struct)

	if isStruct {
		return false
	}

	if len(s.Value.Users()) == 0 {
		return false
	}

	switch instr := s.Value.(type) {
	case *ssa.BinaryOp:
		return !instr.Op.IsComparison()
	}

	return true
}

// String returns the underlying SSA value in human-readable form.
func (s *Step) String() string {
	return s.Value.String()
}