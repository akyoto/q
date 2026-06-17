package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCall(step *Step, instr *ssa.Call) {
	f.moveValuesToRegisters(instr.Arguments, f.CPU.Call.In)
	functionPointer, isPointer := f.ValueToStep[instr.Func]

	if isPointer {
		f.Assembler.Append(&asm.CallRegister{Address: functionPointer.Register})
	} else {
		f.Assembler.Append(&asm.Call{Label: instr.Func.String()})
	}

	if step.Register == -1 || step.Register == f.CPU.Call.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.Call.Out[0],
	})
}