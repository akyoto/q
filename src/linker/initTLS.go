package linker

import (
	"bytes"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/x86"
)

// initTLS initializes the thread-local storage.
func initTLS(program *asm.Assembler, env *core.Environment) {
	tls := ""

	for global := range env.Globals() {
		label := global.File.Package + "." + global.Name
		data := bytes.Repeat([]byte{0}, global.Typ.Size())
		program.Data.SetMutable(label, data)

		// TODO: Make it deterministic.
		if global.ThreadLocal && tls == "" {
			tls = label
		}
	}

	switch env.Build.OS {
	case config.Linux:
		switch env.Build.Arch {
		case config.ARM:
			program.Append(&asm.MoveLabel{
				Destination: arm.X0,
				Label:       tls,
			})

			program.Append(&asm.WriteSystemRegister{
				SystemRegister: arm.TPIDR_EL0,
				Source:         arm.X0,
			})

		case config.X86:
			program.Append(&asm.MoveLabel{
				Destination: x86.R0,
				Label:       tls,
			})

			program.Append(&asm.WriteSystemRegister{
				SystemRegister: x86.FS,
				Source:         x86.R0,
			})
		}
	}
}