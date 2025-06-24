package linker

import (
	"os"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/cli/q/src/elf"
)

// WriteFile writes an executable file to disk.
func WriteFile(executable string, b *build.Build, env *core.Environment) error {
	file, err := os.Create(executable)

	if err != nil {
		return err
	}

	init := env.Functions["core.init"]
	traversed := make(map[*core.Function]bool, len(env.Functions))

	final := asm.Assembler{
		Instructions: make([]asm.Instruction, 0, 8),
		Data:         make(data.Data, 32),
	}

	// This will place the main function immediately after the entry point
	// and also add everything the main function calls recursively.
	init.EachDependency(traversed, func(f *core.Function) {
		final.Merge(&f.Assembler)
	})

	code, data := final.Compile(b)

	switch b.OS {
	case build.Linux:
		elf.Write(file, b, code, data)
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return os.Chmod(executable, 0755)
}