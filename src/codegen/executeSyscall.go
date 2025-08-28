package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeSyscall(step *Step, instr *ssa.Syscall) {
	for i, arg := range instr.Arguments {
		if f.ValueToStep[arg].Register != f.CPU.Syscall.In[i] {
			f.Assembler.Append(&asm.Move{
				Destination: f.CPU.Syscall.In[i],
				Source:      f.ValueToStep[arg].Register,
			})
		}
	}

	f.Assembler.Append(&asm.Syscall{})

	if step.Register == -1 || step.Register == f.CPU.Syscall.Out[0] {
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: step.Register,
		Source:      f.CPU.Syscall.Out[0],
	})
}