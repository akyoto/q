package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeData(step *Step, instr *ssa.Data) {
	f.Assembler.Append(&asm.MoveLabel{
		Destination: step.Register,
		Label:       instr.Label,
	})
}