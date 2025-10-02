package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// fixRegisterConflicts checks for conflicts where 2 values that are live at the same time use the same register.
// It then assigns a new register to the value that was defined earlier.
func (f *Function) fixRegisterConflicts() {
	for _, step := range f.Steps {
		var clobbered []cpu.Register

		switch instr := step.Value.(type) {
		case *ssa.BinaryOp:
			switch instr.Op {
			case token.Div, token.Mod:
				clobbered = f.CPU.DivisionClobbered
			case token.Shl, token.Shr:
				clobbered = f.CPU.ShiftRestricted

				if slices.Contains(f.CPU.ShiftRestricted, step.Register) {
					f.assignFreeRegister(step)
				}
			}

			if step.Register != -1 {
				right := f.ValueToStep[instr.Right]

				if step.Register == right.Register {
					f.assignFreeRegister(right)
				}

				left := f.ValueToStep[instr.Left]

				if instr.Op == token.Mod && step.Register == left.Register {
					f.assignFreeRegister(left)
				}
			}
		case *ssa.Call:
			clobbered = f.CPU.Call.Clobbered
		case *ssa.CallExtern:
			clobbered = f.CPU.ExternCall.Clobbered
		case *ssa.Register:
			if f.build.Arch == config.ARM && step.Register == arm.SP {
				f.assignFreeRegister(step)
			}
		case *ssa.Syscall:
			clobbered = f.CPU.Syscall.Clobbered
		}

		for i, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			if live.Value != step.Value && slices.Contains(clobbered, live.Register) {
				f.assignFreeRegister(live)
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
					f.assignFreeRegister(previous)
				} else {
					f.assignFreeRegister(live)
					break
				}
			}
		}
	}
}