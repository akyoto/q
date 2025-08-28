package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeFromTuple(step *Step, instr *ssa.FromTuple) {
	source := f.CPU.Call.Out[instr.Index]

	if step.Register == source {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      source,
	})
}