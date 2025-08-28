package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeBool(step *Step, instr *ssa.Bool) {
	if step.Register == -1 {
		return
	}

	number := 0

	if instr.Bool {
		number = 1
	}

	f.Assembler.Append(&asm.MoveNumber{
		Destination: step.Register,
		Number:      number,
	})
}