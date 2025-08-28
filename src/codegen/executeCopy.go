package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCopy(step *Step, instr *ssa.Copy) {
	if step.Register == -1 {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.ValueToStep[instr.Value].Register,
	})
}