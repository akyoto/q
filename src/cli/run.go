package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/cli/q/src/memfile"
	"git.urbach.dev/cli/q/src/verbose"
)

// run builds and runs the executable.
func run(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.Show(env)
	file, err := memfile.New(b.Executable())

	if err != nil {
		return exit(err)
	}

	linker.Write(file, env)
	err = memfile.Exec(file)

	if err != nil {
		return exit(err)
	}

	return 0
}