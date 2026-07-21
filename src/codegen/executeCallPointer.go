package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCallPointer(step *Step, instr *ssa.CallPointer) {
	f.moveValuesToRegisters(instr.Arguments[1:], f.CPU.Call.In)
	functionPointer := f.ValueToStep[instr.Arguments[0]]
	address := f.resolveOperand(functionPointer, step.Live)

	f.Assembler.Append(&asm.CallRegister{
		Address: address,
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