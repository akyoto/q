package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeParameter(step *Step, instr *ssa.Parameter) {
	source := f.CPU.Call.In[instr.Index]

	if step.Register == source {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      source,
	})
}