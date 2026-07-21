package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCall(step *Step, instr *ssa.Call) {
	f.moveValuesToRegisters(instr.Arguments, f.CPU.Call.In)

	f.Assembler.Append(&asm.Call{
		Label: instr.Func.String(),
	})

	destination := step.Register

	if destination == -1 || destination == f.CPU.Call.Out[0] {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpill(step, f.CPU.Call.Out[0])
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      f.CPU.Call.Out[0],
	})
}