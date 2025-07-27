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
			switch instr.Op {
			case token.Div, token.Mod:
				clobbered = f.CPU.Division
			case token.Shl, token.Shr:
				clobbered = f.CPU.Shift

				if slices.Contains(f.CPU.Shift, step.Register) {
					users := step.Value.Users()

					if len(users) > 0 {
						alive := f.ValueToStep[users[len(users)-1]].Index
						step.Register = f.findFreeRegister(f.Steps[step.Index:alive])
					}
				}
			}

			right := f.ValueToStep[instr.Right]

			if step.Register == right.Register {
				right.Register = f.findFreeRegister(f.Steps[right.Index:stepIndex])
			}

			left := f.ValueToStep[instr.Left]

			if instr.Op == token.Mod && step.Register == left.Register {
				left.Register = f.findFreeRegister(f.Steps[left.Index:stepIndex])
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
					f.assignFreeRegister(previous)
				} else {
					f.assignFreeRegister(live)
					break
				}
			}
		}
	}
}