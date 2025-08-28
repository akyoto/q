package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCallExtern(step *Step, instr *ssa.CallExtern) {
	args := instr.Arguments
	var pushed []cpu.Register

	for i, arg := range slices.Backward(args) {
		if i >= len(f.CPU.ExternCall.In) {
			pushed = append(pushed, f.ValueToStep[arg].Register)
			continue
		}

		if f.ValueToStep[arg].Register == f.CPU.ExternCall.In[i] {
			continue
		}

		f.Assembler.Append(&asm.Move{
			Destination: f.CPU.ExternCall.In[i],
			Source:      f.ValueToStep[arg].Register,
		})
	}

	// Pushing an odd number of registers would not maintain the 16-byte
	// stack alignment, so we allocate additional 8 bytes before pushing
	// the 5th argument.
	if len(pushed)&1 != 0 {
		f.Assembler.Append(&asm.SubtractNumber{Destination: f.CPU.StackPointer, Source: f.CPU.StackPointer, Number: 8})
	}

	// TODO: Replace push instructions with store instructions using a fixed offset.
	if len(pushed) > 0 {
		f.Assembler.Append(&asm.Push{Registers: pushed})
	}

	f.Assembler.Append(&asm.CallExtern{Library: instr.Func.Package, Function: instr.Func.Name})

	if len(pushed) > 0 {
		f.Assembler.Append(&asm.Pop{Registers: pushed})
	}

	if len(pushed)&1 != 0 {
		f.Assembler.Append(&asm.AddNumber{Destination: f.CPU.StackPointer, Source: f.CPU.StackPointer, Number: 8})
	}

	if step.Register == -1 || step.Register == f.CPU.ExternCall.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.ExternCall.Out[0],
	})
}