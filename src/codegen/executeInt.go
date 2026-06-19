package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeInt(step *Step, instr *ssa.Int) {
	if step.Register == -1 {
		return
	}

	if f.isSpilled(step.Register) {
		f.storeSpillNumber(step, instr)
		return
	}

	f.Assembler.Append(&asm.MoveNumber{
		Destination: step.Register,
		Number:      instr.Int,
	})
}