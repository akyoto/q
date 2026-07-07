package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/verbose"
)

// functions shows the list of functions that are used in a build.
func functions(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	b.Dry = true
	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.Functions(env)
	return success
}