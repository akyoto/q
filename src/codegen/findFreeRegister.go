package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// findFreeRegister finds a free register for the given value.
func (f *Function) findFreeRegister(value *Step) cpu.Register {
	usedRegisters := 0

	for _, step := range f.Steps {
		// These checks need to happen regardless of whether the value is alive after execution.
		// If it is used as an operand, the operand restrictions of the architecture apply.
		switch instr := step.Value.(type) {
		case *ssa.BinaryOp:
			switch instr.Op {
			case token.Div, token.Mod:
				if instr.Right == value.Value {
					for _, reg := range f.CPU.DivisorRestricted {
						usedRegisters |= (1 << reg)
					}
				}

			case token.Shl, token.Shr:
				if instr.Left == value.Value {
					for _, reg := range f.CPU.ShiftRestricted {
						usedRegisters |= (1 << reg)
					}
				}
			}
		}

		// If it's not alive in this step, ignore it.
		if !slices.Contains(step.Live, value) {
			continue
		}

		// Mark all the neighbor registers that are alive
		// at the same time as used.
		for _, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			usedRegisters |= (1 << live.Register)
		}

		// Ignore the definition itself.
		if step == value {
			continue
		}

		// Find all the registers that this instruction
		// would clobber and mark them as used.
		var clobbered []cpu.Register

		switch instr := step.Value.(type) {
		case *ssa.BinaryOp:
			switch instr.Op {
			case token.Div, token.Mod:
				clobbered = f.CPU.DivisionClobbered
			case token.Shl, token.Shr:
				clobbered = f.CPU.ShiftClobbered
			}
		case *ssa.Call:
			clobbered = f.CPU.Call.Clobbered
		case *ssa.CallExtern:
			clobbered = f.CPU.ExternCall.Clobbered
		case *ssa.Parameter:
			usedRegisters |= (1 << f.CPU.Call.In[instr.Index])
		case *ssa.Syscall:
			clobbered = f.CPU.Syscall.Clobbered
		}

		for _, reg := range clobbered {
			usedRegisters |= (1 << reg)
		}
	}

	// Pick one of the register hints if possible.
	for _, reg := range value.Hints {
		if usedRegisters&(1<<reg) == 0 {
			return reg
		}
	}

	// Pick a general purpose register that's not used yet.
	for _, reg := range f.CPU.General {
		if usedRegisters&(1<<reg) == 0 {
			return reg
		}
	}

	panic("no free registers")
}