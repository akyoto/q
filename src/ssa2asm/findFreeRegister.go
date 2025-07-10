package ssa2asm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// findFreeRegister finds a free register within the given slice of steps.
func (f *Compiler) findFreeRegister(steps []Step) cpu.Register {
	usedRegisters := 0

	for _, step := range steps {
		for _, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			usedRegisters |= (1 << live.Register)
		}
	}

	for _, candidate := range f.CPU.General {
		if usedRegisters&(1<<candidate) == 0 {
			return candidate
		}
	}

	panic("no free registers")
}