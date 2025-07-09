package linker

import (
	"io"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/build"
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

	init := env.Functions["run.init"]
	traversed := make(map[*core.Function]bool, len(env.Functions))

	// This will place the init function immediately after the entry point
	// and also add everything the init function calls recursively.
	init.EachDependency(traversed, func(f *core.Function) {
		program.Merge(&f.Assembler)
	})

	code, data, libs := program.Compile(env.Build)

	switch env.Build.OS {
	case build.Linux:
		elf.Write(writer, env.Build, code, data)
	case build.Mac:
		macho.Write(writer, env.Build, code, data)
	case build.Windows:
		pe.Write(writer, env.Build, code, data, libs)
	}
}