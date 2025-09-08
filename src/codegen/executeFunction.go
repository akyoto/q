package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeFunction(step *Step, instr *ssa.Function) {
	if step.Register == -1 {
		return
	}

	f.Assembler.Append(&asm.MoveLabel{
		Destination: step.Register,
		Label:       instr.String(),
	})
}