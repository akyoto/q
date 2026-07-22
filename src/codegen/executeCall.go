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

	f.moveCallResult(step, f.CPU.Call.Out[0])
}