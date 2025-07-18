package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// fixRegisterConflicts checks for conflicts where 2 values that are live at the same time use the same register.
// It then assigns a new register to the value that was defined earlier.
func (f *Function) fixRegisterConflicts() {
	for stepIndex, step := range f.Steps {
		var clobbered []cpu.Register

		switch instr := step.Value.(type) {
		case *ssa.BinaryOp:
			if instr.Op == token.Div {
				clobbered = f.CPU.Division
			}

			right := f.ValueToStep[instr.Right]

			if step.Register == right.Register {
				right.Register = f.findFreeRegister(f.Steps[right.Index:stepIndex])
			}
		case *ssa.Call:
			clobbered = f.CPU.Call.Clobbered
		case *ssa.CallExtern:
			clobbered = f.CPU.ExternCall.Clobbered
		case *ssa.Syscall:
			clobbered = f.CPU.Syscall.Clobbered
		}

		for i, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			if live.Value != step.Value && slices.Contains(clobbered, live.Register) {
				live.Register = f.findFreeRegister(f.Steps[live.Index : stepIndex+1])
				continue
			}

			for _, previous := range step.Live[:i] {
				if previous.Register == -1 {
					continue
				}

				if previous.Register != live.Register {
					continue
				}

				if previous.Index < live.Index {
					previous.Register = f.findFreeRegister(f.Steps[previous.Index : stepIndex+1])
				} else {
					live.Register = f.findFreeRegister(f.Steps[live.Index : stepIndex+1])
					break
				}
			}
		}
	}
}