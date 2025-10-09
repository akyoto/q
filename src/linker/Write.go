package linker

import (
	"bytes"
	"io"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/macho"
	"git.urbach.dev/cli/q/src/pe"
)

// Write writes an executable to the given writer.
func Write(writer io.WriteSeeker, env *core.Environment) {
	program := asm.Assembler{
		Instructions: make([]asm.Instruction, 0, 256),
		Data: data.Data{
			Immutable: make(map[string][]byte, 16),
			Mutable:   make(map[string][]byte, 16),
		},
	}

	for f := range env.LiveFunctions() {
		program.Merge(&f.Assembler)
	}

	for global := range env.Globals() {
		label := global.File.Package + "." + global.Name
		data := bytes.Repeat([]byte{0}, global.Typ.Size())
		program.Data.SetMutable(label, data)
	}

	build := env.Build
	code, data, libs := program.Compile(build)

	switch build.OS {
	case config.Linux:
		elf.Write(writer, build, code, data)
	case config.Mac:
		macho.Write(writer, build, code, data)
	case config.Windows:
		pe.Write(writer, build, code, data, libs)
	}
}