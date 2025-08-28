package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCopy(step *Step, instr *ssa.Copy) {
	if step.Register == -1 {
		return
	}

	copy := f.ValueToStep[instr.Value]

	if step.Register == copy.Register {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      copy.Register,
	})
}