package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

func (f *Function) executeBool(step *Step, instr *ssa.Bool) {
	destination := step.Register

	if destination == -1 {
		return
	}

	number := 0

	if instr.Bool {
		number = 1
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpillNumber(step, types.Bool, number)
		return
	}

	f.Assembler.Append(&asm.MoveNumber{
		Destination: destination,
		Number:      number,
	})
}