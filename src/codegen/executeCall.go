package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCall(step *Step, instr *ssa.Call) {
	args := instr.Arguments

	for i, arg := range args {
		if f.ValueToStep[arg].Register == f.CPU.Call.In[i] {
			continue
		}

		f.Assembler.Append(&asm.Move{
			Destination: f.CPU.Call.In[i],
			Source:      f.ValueToStep[arg].Register,
		})
	}

	f.Assembler.Append(&asm.Call{Label: instr.Func.String()})

	if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.Call.Out[0],
	})
}