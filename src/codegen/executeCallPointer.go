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

	f.moveCallResult(step, f.CPU.Call.Out[0])
}