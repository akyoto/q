package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeSyscall(step *Step, instr *ssa.Syscall) {
	f.moveValuesToRegisters(instr.Arguments, f.CPU.Syscall.In)
	f.Assembler.Append(&asm.Syscall{})
	f.moveCallResult(step, f.CPU.Syscall.Out[0])
}