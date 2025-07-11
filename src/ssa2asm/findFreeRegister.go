package ssa2asm

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// findFreeRegister finds a free register within the given slice of steps.
func (f *Compiler) findFreeRegister(steps []*Step) cpu.Register {
	usedRegisters := 0

	for _, step := range steps {
		for _, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			usedRegisters |= (1 << live.Register)
		}

		var volatileRegisters []cpu.Register

		switch instr := step.Value.(type) {
		case *ssa.Call:
			volatileRegisters = f.CPU.Call.Volatile
		case *ssa.CallExtern:
			volatileRegisters = f.CPU.ExternCall.Volatile
		case *ssa.Parameter:
			usedRegisters |= (1 << f.CPU.Call.In[instr.Index])
		case *ssa.Syscall:
			volatileRegisters = f.CPU.Syscall.Volatile
		}

		for _, volatile := range volatileRegisters {
			usedRegisters |= (1 << volatile)
		}
	}

	for _, candidate := range f.CPU.General {
		if usedRegisters&(1<<candidate) == 0 {
			return candidate
		}
	}

	panic("no free registers")
}