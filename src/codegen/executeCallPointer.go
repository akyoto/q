package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCallPointer(step *Step, instr *ssa.CallPointer) {
	f.moveValuesToRegisters(instr.Arguments[1:], f.CPU.Call.In)
	functionPointer := f.ValueToStep[instr.Arguments[0]]

	f.Assembler.Append(&asm.CallRegister{
		Address: functionPointer.Register,
	})

	if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.Call.Out[0],
	})
}