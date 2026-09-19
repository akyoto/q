package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeGlobal(step *Step, instr *ssa.Global) {
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	if instr.ThreadLocal {
		f.arch.loadTLS(f, destination, instr.Label)
	} else {
		f.Assembler.Append(&asm.MoveLabel{
			Destination: destination,
			Label:       instr.Label,
		})
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}