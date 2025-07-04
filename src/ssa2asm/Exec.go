package ssa2asm

import (
	"strings"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Compiler) Exec(step *Step) {
	switch instr := step.Value.(type) {
	case *ssa.BinaryOp:
		f.Assembler.Append(&asm.AddRegisterRegister{
			Destination: step.Register,
			Source:      f.ValueToStep[instr.Left].Register,
			Operand:     f.ValueToStep[instr.Right].Register,
		})

	case *ssa.Bytes:
		f.Count.Data++
		label := f.CreateLabel("data", f.Count.Data)
		f.Assembler.SetData(label.Name, instr.Bytes)

		f.Assembler.Append(&asm.MoveRegisterLabel{
			Destination: step.Register,
			Label:       label.Name,
		})

	case *ssa.Call:
		args := instr.Arguments[1:]
		offset := 0

		for i, arg := range args {
			structure, isStruct := args[i].(*ssa.Struct)

			if isStruct {
				for _, field := range structure.Arguments {
					if f.ValueToStep[field].Register != f.CPU.Call[offset+i] {
						f.Assembler.Append(&asm.MoveRegisterRegister{
							Destination: f.CPU.Call[offset+i],
							Source:      f.ValueToStep[field].Register,
						})
					}

					offset++
				}

				offset--
			} else {
				if f.ValueToStep[arg].Register != f.CPU.Call[offset+i] {
					f.Assembler.Append(&asm.MoveRegisterRegister{
						Destination: f.CPU.Call[offset+i],
						Source:      f.ValueToStep[arg].Register,
					})
				}
			}
		}

		fn := instr.Arguments[0].(*ssa.Function)
		f.Assembler.Append(&asm.Call{Label: fn.UniqueName})

		if step.Register == -1 || step.Register == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Return[0],
		})

	case *ssa.CallExtern:
		args := instr.Arguments[1:]

		for i, arg := range args {
			if f.ValueToStep[arg].Register != f.CPU.ExternCall[i] {
				f.Assembler.Append(&asm.MoveRegisterRegister{
					Destination: f.CPU.ExternCall[i],
					Source:      f.ValueToStep[arg].Register,
				})
			}
		}

		fn := instr.Arguments[0].(*ssa.Function)
		dot := strings.IndexByte(fn.UniqueName, '.')
		library := fn.UniqueName[:dot]
		function := fn.UniqueName[dot+1:]
		f.Assembler.Append(&asm.CallExtern{Library: library, Function: function})

		if step.Register == -1 || step.Register == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Return[0],
		})

	case *ssa.Int:
		f.Assembler.Append(&asm.MoveRegisterNumber{
			Destination: step.Register,
			Number:      instr.Int,
		})

	case *ssa.Parameter:
		source := f.CPU.Call[instr.Index]

		if step.Register == -1 || step.Register == source {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      source,
		})

	case *ssa.Field:
		parameter := instr.Object.(*ssa.Parameter)
		field := instr.Field
		source := f.CPU.Call[parameter.Index+field.Index]

		if step.Register == -1 || step.Register == source {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      source,
		})

	case *ssa.Syscall:
		for i, arg := range instr.Arguments {
			if f.ValueToStep[arg].Register != f.CPU.Syscall[i] {
				f.Assembler.Append(&asm.MoveRegisterRegister{
					Destination: f.CPU.Syscall[i],
					Source:      f.ValueToStep[arg].Register,
				})
			}
		}

		f.Assembler.Append(&asm.Syscall{})

		if step.Register == -1 || step.Register == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Return[0],
		})
	}
}