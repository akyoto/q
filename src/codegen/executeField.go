package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeField(step *Step, instr *ssa.Field) {
	source := f.CPU.Call.Out[instr.Index]
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if destination == source {
		return
	}

	if isSpilled {
		f.storeSpill(step, source)
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      source,
	})
}