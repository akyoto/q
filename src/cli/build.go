package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/cli/q/src/verbose"
)

// build parses the arguments and creates a build.
func build(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.Show(env)

	if b.Dry {
		return 0
	}

	err = linker.WriteFile(b.Executable(), env)
	return exit(err)
}