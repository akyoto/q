package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/x86"
)

const (
	WindowsTLSOffset = 0x480
	WindowsTLSSize   = 0x200
)

func (f *Function) executeGlobal(step *Step, instr *ssa.Global) {
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	if instr.ThreadLocal {
		switch f.build.OS {
		case config.Linux:
			switch f.build.Arch {
			case config.ARM:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    destination,
					SystemRegister: arm.TPIDR_EL0,
				})

			case config.X86:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    destination,
					SystemRegister: x86.FS,
				})
			}
		case config.Windows:
			switch f.build.Arch {
			case config.ARM:
				f.Assembler.Append(&asm.AddNumber{
					Destination: destination,
					Source:      arm.X18,
					Number:      0x1000,
				})

				f.Assembler.Append(&asm.AddNumber{
					Destination: destination,
					Source:      destination,
					Number:      WindowsTLSOffset + WindowsTLSSize - 0x20,
				})

			case config.X86:
				f.Assembler.Append(&asm.ReadSystemRegister{
					Destination:    destination,
					SystemRegister: x86.GS,
				})

				f.Assembler.Append(&asm.AddNumber{
					Destination: destination,
					Source:      destination,
					Number:      0x1000 + WindowsTLSOffset + WindowsTLSSize - 0x20,
				})
			}

		default:
			f.Assembler.Append(&asm.MoveLabel{
				Destination: destination,
				Label:       instr.Label,
			})
		}
	} else {
		f.Assembler.Append(&asm.MoveLabel{
			Destination: destination,
			Label:       instr.Label,
		})
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}