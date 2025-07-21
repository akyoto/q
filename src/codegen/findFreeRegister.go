package codegen

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// findFreeRegister finds a free register within the given slice of steps.
func (f *Function) findFreeRegister(steps []*step) cpu.Register {
	usedRegisters := 0

	for i, step := range steps {
		for _, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			usedRegisters |= (1 << live.Register)
		}

		if i == 0 {
			continue
		}

		var volatileRegisters []cpu.Register

		switch instr := step.Value.(type) {
		case *ssa.BinaryOp:
			switch instr.Op {
			case token.Div, token.Mod:
				volatileRegisters = f.CPU.Division
			case token.Shl, token.Shr:
				volatileRegisters = f.CPU.Shift
			}
		case *ssa.Call:
			volatileRegisters = f.CPU.Call.Clobbered
		case *ssa.CallExtern:
			volatileRegisters = f.CPU.ExternCall.Clobbered
		case *ssa.Parameter:
			usedRegisters |= (1 << f.CPU.Call.In[instr.Index])
		case *ssa.Syscall:
			volatileRegisters = f.CPU.Syscall.Clobbered
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