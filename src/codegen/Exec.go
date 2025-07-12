package codegen

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
		args := instr.Arguments

		for i, arg := range args {
			if f.ValueToStep[arg].Register == f.CPU.Call.In[i] {
				continue
			}

			f.Assembler.Append(&asm.MoveRegisterRegister{
				Destination: f.CPU.Call.In[i],
				Source:      f.ValueToStep[arg].Register,
			})
		}

		f.Assembler.Append(&asm.Call{Label: instr.Func.UniqueName})

		if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Call.Out[0],
		})

	case *ssa.CallExtern:
		f.Assembler.Append(&asm.CallExternStart{})
		args := instr.Arguments

		for i, arg := range args {
			if i >= len(f.CPU.ExternCall.In) {
				f.Assembler.Append(&asm.PushRegister{
					Register: f.ValueToStep[arg].Register,
				})

				continue
			}

			if f.ValueToStep[arg].Register == f.CPU.ExternCall.In[i] {
				continue
			}

			f.Assembler.Append(&asm.MoveRegisterRegister{
				Destination: f.CPU.ExternCall.In[i],
				Source:      f.ValueToStep[arg].Register,
			})
		}

		dot := strings.IndexByte(instr.Func.UniqueName, '.')
		library := instr.Func.UniqueName[:dot]
		function := instr.Func.UniqueName[dot+1:]
		f.Assembler.Append(&asm.CallExtern{Library: library, Function: function})
		f.Assembler.Append(&asm.CallExternEnd{})

		if step.Register == -1 || step.Register == f.CPU.ExternCall.Out[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.ExternCall.Out[0],
		})

	case *ssa.Int:
		f.Assembler.Append(&asm.MoveRegisterNumber{
			Destination: step.Register,
			Number:      instr.Int,
		})

	case *ssa.Parameter:
		source := f.CPU.Call.In[instr.Index]

		if step.Register == -1 || step.Register == source {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      source,
		})

	case *ssa.Field:
		structure := instr.Object.(*ssa.Struct)
		field := structure.Arguments[instr.Field.Index]
		source := f.ValueToStep[field].Register

		if step.Register == -1 || step.Register == source {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      source,
		})

	case *ssa.Syscall:
		for i, arg := range instr.Arguments {
			if f.ValueToStep[arg].Register != f.CPU.Syscall.In[i] {
				f.Assembler.Append(&asm.MoveRegisterRegister{
					Destination: f.CPU.Syscall.In[i],
					Source:      f.ValueToStep[arg].Register,
				})
			}
		}

		f.Assembler.Append(&asm.Syscall{})

		if step.Register == -1 || step.Register == f.CPU.Syscall.Out[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Syscall.Out[0],
		})
	}
}