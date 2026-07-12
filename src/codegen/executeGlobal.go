package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/x86"
)

const (
	WindowsTLSOffset = 0x1480
	WindowsTLSSize   = 0x200
)

func (f *Function) executeGlobal(step *Step, instr *ssa.Global) {
	if instr.ThreadLocal {
		switch f.build.OS {
		case config.Linux:
			switch f.build.Arch {
			case config.ARM:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    step.Register,
					SystemRegister: arm.TPIDR_EL0,
				})

			case config.X86:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    step.Register,
					SystemRegister: x86.FS,
				})
			}
		case config.Windows:
			switch f.build.Arch {
			case config.ARM:
				f.Assembler.Append(&asm.Move{
					Destination: step.Register,
					Source:      arm.X18,
				})

			case config.X86:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    step.Register,
					SystemRegister: x86.GS,
				})
			}

			f.Assembler.Append(&asm.AddNumber{
				Destination: step.Register,
				Source:      step.Register,
				Number:      WindowsTLSOffset + WindowsTLSSize - 0x20,
			})
		default:
			f.Assembler.Append(&asm.MoveLabel{
				Destination: step.Register,
				Label:       instr.Label,
			})
		}
	} else {
		f.Assembler.Append(&asm.MoveLabel{
			Destination: step.Register,
			Label:       instr.Label,
		})
	}
}