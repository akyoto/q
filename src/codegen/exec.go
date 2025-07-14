package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// exec executes an instruction.
func (f *Function) exec(step *step) {
	switch instr := step.Value.(type) {
	case *ssa.BinaryOp:
		left := f.ValueToStep[instr.Left]
		right := f.ValueToStep[instr.Right]

		switch instr.Op {
		case token.Add:
			f.Assembler.Append(&asm.AddRegisterRegister{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Div:
			f.Assembler.Append(&asm.DivRegisterRegister{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Mul:
			f.Assembler.Append(&asm.MulRegisterRegister{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		case token.Sub:
			f.Assembler.Append(&asm.SubRegisterRegister{
				Destination: step.Register,
				Source:      left.Register,
				Operand:     right.Register,
			})

		default:
			panic("not implemented")
		}

	case *ssa.Bytes:
		f.Count.Data++
		label := f.createLabel("data", f.Count.Data)
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

		f.Assembler.Append(&asm.Call{Label: instr.Func.String()})

		if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: step.Register,
			Source:      f.CPU.Call.Out[0],
		})

	case *ssa.CallExtern:
		args := instr.Arguments
		var pushed []cpu.Register

		for i, arg := range args {
			if i >= len(f.CPU.ExternCall.In) {
				pushed = append(pushed, f.ValueToStep[arg].Register)
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

		if len(pushed) > 0 {
			f.Assembler.Append(&asm.PushRegisters{Registers: pushed})
		}

		f.Assembler.Append(&asm.CallExtern{Library: instr.Func.Package, Function: instr.Func.Name})

		if len(pushed) > 0 {
			f.Assembler.Append(&asm.PopRegisters{Registers: pushed})
		}

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