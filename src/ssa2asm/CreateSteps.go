package ssa2asm

import (
	"slices"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Compiler) CreateSteps(ir ssa.IR) []Step {
	count := ir.CountValues()
	steps := make([]Step, count)
	f.ValueToStep = make(map[ssa.Value]*Step, count)

	for i, instr := range ir.Values {
		steps[i].Index = i
		steps[i].Value = instr
		steps[i].Register = -1
		f.ValueToStep[instr] = &steps[i]
	}

	for i, instr := range ir.Values {
		switch instr := instr.(type) {
		case *ssa.Call:
			for paramIndex, param := range instr.Arguments {
				f.ValueToStep[param].Hint(f.CPU.Call.In[paramIndex])
			}

		case *ssa.CallExtern:
			for r, param := range instr.Arguments {
				if r >= len(f.CPU.ExternCall.In) {
					// Temporary hack to allow arguments 5 and 6 to be hinted as r10 and r11, then pushed later.
					f.ValueToStep[param].Hint(f.CPU.ExternCall.Volatile[1+r])
					continue
				}

				f.ValueToStep[param].Hint(f.CPU.ExternCall.In[r])
			}

		case *ssa.Parameter:
			f.ValueToStep[instr].Register = f.CPU.Call.In[instr.Index]

		case *ssa.Return:
			for r, param := range instr.Arguments {
				f.ValueToStep[param].Hint(f.CPU.Call.Out[r])
			}

		case *ssa.Syscall:
			for r, param := range instr.Arguments {
				f.ValueToStep[param].Hint(f.CPU.Syscall.In[r])
			}
		}

		users := instr.Users()

		if len(users) == 0 {
			continue
		}

		liveStart := i
		_, isParam := instr.(*ssa.Parameter)

		if isParam {
			liveStart = 0
		}

		liveEnd := f.ValueToStep[users[len(users)-1]].Index
		instrStep := f.ValueToStep[instr]

		for live := liveStart; live < liveEnd; live++ {
			steps[live].Live = append(steps[live].Live, instrStep)
		}
	}

	for stepIndex, step := range steps {
		for i, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			var (
				oldRegister       = cpu.Register(-1)
				volatileRegisters []cpu.Register
			)

			switch instr := step.Value.(type) {
			case *ssa.Call:
				volatileRegisters = f.CPU.Call.Volatile
			case *ssa.CallExtern:
				volatileRegisters = f.CPU.ExternCall.Volatile
			case *ssa.Parameter:
				oldRegister = f.CPU.Call.In[instr.Index]
			case *ssa.Syscall:
				volatileRegisters = f.CPU.Syscall.Volatile
			}

			if live.Value != step.Value && slices.Contains(volatileRegisters, live.Register) {
				live.Register = f.findFreeRegister(steps[live.Index : stepIndex+1])
				goto next
			}

			for _, previous := range step.Live[:i] {
				if previous.Register == -1 {
					continue
				}

				if previous.Register != live.Register && previous.Register != oldRegister {
					continue
				}

				if previous.Index < live.Index {
					previous.Register = f.findFreeRegister(steps[previous.Index : stepIndex+1])
				} else {
					live.Register = f.findFreeRegister(steps[live.Index : stepIndex+1])
					goto next
				}
			}
		next:
		}
	}

	return steps
}