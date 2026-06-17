package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCall(step *Step, instr *ssa.Call) {
	args := instr.Arguments

	for i, arg := range args {
		sourceStep := f.ValueToStep[arg]
		source := sourceStep.Register
		destination := f.CPU.Call.In[i]

		if f.isSpilled(source) {
			f.loadSpill(sourceStep, destination)
			continue
		}

		if source == destination {
			continue
		}

		f.Assembler.Append(&asm.Move{
			Destination: destination,
			Source:      source,
		})
	}

	functionPointer, isPointer := f.ValueToStep[instr.Func]

	if isPointer {
		f.Assembler.Append(&asm.CallRegister{Address: functionPointer.Register})
	} else {
		f.Assembler.Append(&asm.Call{Label: instr.Func.String()})
	}

	if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.Call.Out[0],
	})
}