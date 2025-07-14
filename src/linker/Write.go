package linker

import (
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
		Instructions: make([]asm.Instruction, 0, 32),
		Data:         make(data.Data, 32),
	}

	init := env.Init
	traversed := make(map[*core.Function]bool, env.NumFunctions)

	// This will place the init function immediately after the entry point
	// and also add everything the init function calls recursively.
	init.EachDependency(traversed, func(f *core.Function) {
		program.Merge(&f.Assembler)
	})

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