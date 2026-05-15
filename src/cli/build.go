package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
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

	if b.Dry {
		return success
	}

	err = linker.WriteFile(b.Executable(), env)
	return exit(err)
}