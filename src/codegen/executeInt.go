package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeInt(step *Step, instr *ssa.Int) {
	destination := step.Register

	if destination == -1 {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpillNumber(step, instr.Type(), instr.Int)
		return
	}

	f.Assembler.Append(&asm.MoveNumber{
		Destination: destination,
		Number:      instr.Int,
	})
}