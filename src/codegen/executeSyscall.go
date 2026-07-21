package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeSyscall(step *Step, instr *ssa.Syscall) {
	f.moveValuesToRegisters(instr.Arguments, f.CPU.Syscall.In)
	f.Assembler.Append(&asm.Syscall{})
	destination := step.Register

	if destination == -1 || destination == f.CPU.Syscall.Out[0] {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpill(step, f.CPU.Syscall.Out[0])
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      f.CPU.Syscall.Out[0],
	})
}