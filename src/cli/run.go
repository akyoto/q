package cli

import (
	"fmt"
	"os"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/cli/q/src/memfile"
)

// run builds and runs the executable.
func run(args []string) int {
	b, err := newBuildFromArgs(args)

	if err != nil {
		return exit(err)
	}

	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	file, err := memfile.New(b.Executable())

	if err != nil {
		return exit(err)
	}

	linker.Write(file, env)
	err = memfile.Exec(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exit(err)
	}

	return 0
}