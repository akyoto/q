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
		argStep := f.ValueToStep[arg]
		source := f.resolveOperand(argStep, step.Live)

		if i >= len(f.CPU.ExternCall.In) {
			pushed = append(pushed, source)
			continue
		}

		if source == f.CPU.ExternCall.In[i] {
			continue
		}

		f.Assembler.Append(&asm.Move{
			Destination: f.CPU.ExternCall.In[i],
			Source:      source,
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

	f.Assembler.Append(&asm.CallExtern{Library: instr.Func.Package(), Function: instr.Func.Name()})

	if len(pushed) > 0 {
		f.Assembler.Append(&asm.Pop{Registers: pushed})
	}

	if len(pushed)&1 != 0 {
		f.Assembler.Append(&asm.AddNumber{Destination: f.CPU.StackPointer, Source: f.CPU.StackPointer, Number: 8})
	}

	f.moveCallResult(step, f.CPU.ExternCall.Out[0])
}