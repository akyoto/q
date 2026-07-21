package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeParameter(step *Step, instr *ssa.Parameter) {
	destination := step.Register

	if destination == -1 {
		return
	}

	source := f.CPU.Call.In[instr.Index]

	if destination == source {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpill(step, source)
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      source,
	})
}