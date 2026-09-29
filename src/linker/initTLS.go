package linker

import (
	"bytes"
	"encoding/binary"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/x86"
)

// initTLS initializes the thread-local storage.
func initTLS(program *asm.Assembler, env *core.Environment) {
	used := false
	hasTLSSize := false
	base := ""

	for global := range env.Globals() {
		if !global.ThreadLocal {
			if global.Used.Load() == 0 {
				continue
			}

			label := global.File.Package + "." + global.Name
			program.Data.SetMutable(label, bytes.Repeat([]byte{0}, global.Typ.Size()))

			if label == "thread.tlsSize" {
				hasTLSSize = true
			}

			continue
		}

		label := global.File.Package + "." + global.Name
		program.Data.SetTLS(label, bytes.Repeat([]byte{0}, global.Typ.Size()))

		if base == "" || label < base {
			base = label
		}

		if global.Used.Load() > 0 {
			used = true
		}
	}

	if !used {
		return
	}

	if env.Build.OS == config.Linux && hasTLSSize {
		_, tlsSize := env.TLSLayout()
		size := make([]byte, 8)
		binary.LittleEndian.PutUint64(size, uint64(tlsSize))
		program.Data.SetMutable("thread.tlsSize", size)
	}

	switch env.Build.OS {
	case config.Linux:
		switch env.Build.Arch {
		case config.ARM:
			program.Append(&asm.MoveLabel{
				Destination: arm.X0,
				Label:       base,
			})

			program.Append(&asm.WriteSystemRegister{
				SystemRegister: arm.TPIDR_EL0,
				Source:         arm.X0,
			})
		case config.X86:
			program.Append(&asm.MoveLabel{
				Destination: x86.R0,
				Label:       base,
			})

			program.Append(&asm.WriteSystemRegister{
				SystemRegister: x86.FS,
				Source:         x86.R0,
			})
		}
	}
}