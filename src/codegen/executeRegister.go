package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeRegister(step *Step, instr *ssa.Register) {
	if step.Register == instr.Register {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      instr.Register,
	})
}