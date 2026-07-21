package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeRegister(step *Step, instr *ssa.Register) {
	destination := step.Register

	if destination == -1 {
		return
	}

	if destination == instr.Register {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpill(step, instr.Register)
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      instr.Register,
	})
}