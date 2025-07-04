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
			offset := 0

			for r, param := range instr.Arguments[1:] {
				structure, isStruct := param.(*ssa.Struct)

				if isStruct {
					for _, field := range structure.Arguments {
						f.ValueToStep[field].Hint(f.CPU.Call[offset+r])
						offset++
					}

					offset--
				} else {
					f.ValueToStep[param].Hint(f.CPU.Call[offset+r])
				}
			}

		case *ssa.CallExtern:
			for r, param := range instr.Arguments[1:] {
				f.ValueToStep[param].Hint(f.CPU.ExternCall[r])
			}

		case *ssa.Parameter:
			f.ValueToStep[instr].Register = f.CPU.Call[instr.Index]

		case *ssa.Return:
			for r, param := range instr.Arguments {
				f.ValueToStep[param].Hint(f.CPU.Return[r])
			}

		case *ssa.Syscall:
			for r, param := range slices.Backward(instr.Arguments) {
				f.ValueToStep[param].Hint(f.CPU.Syscall[r])
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

	for _, step := range steps {
		for liveIndex, live := range step.Live {
			if live.Register == -1 {
				continue
			}

			oldRegister := cpu.Register(-1)
			liveParam, isParam := live.Value.(*ssa.Parameter)

			if isParam {
				oldRegister = f.CPU.Call[liveParam.Index]
			}

			for _, existing := range step.Live[:liveIndex] {
				if existing.Register == -1 {
					continue
				}

				if existing.Register == live.Register || existing.Register == oldRegister {
					a := existing.Index
					b := live.Index
					freeRegister := cpu.Register(15)

					if a < b {
						existing.Register = freeRegister
					} else {
						live.Register = freeRegister
					}
				}
			}
		}
	}

	return steps
}