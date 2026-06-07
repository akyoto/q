package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeGlobal(step *Step, instr *ssa.Global) {
	f.Assembler.Append(&asm.MoveLabel{
		Destination: step.Register,
		Label:       instr.Label,
	})
}