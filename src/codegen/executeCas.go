package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/x86"
)

func (f *Function) executeCas(step *Step, instr *ssa.Cas) {
	address := f.ValueToStep[instr.Arguments[0]]
	oldValue := f.ValueToStep[instr.Arguments[1]]
	newValue := f.ValueToStep[instr.Arguments[2]]
	oldValueRegister := oldValue.Register

	if f.build.Arch == config.X86 && oldValueRegister != x86.R0 {
		f.Assembler.Append(&asm.Move{
			Destination: x86.R0,
			Source:      oldValueRegister,
		})

		oldValueRegister = x86.R0
	}

	f.Assembler.Append(&asm.CompareAndSwap{
		OldValue: oldValueRegister,
		NewValue: newValue.Register,
		Address:  address.Register,
		Length:   8,
	})
}