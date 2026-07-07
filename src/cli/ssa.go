package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/verbose"
)

// ssa shows the SSA form.
func ssa(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	b.Dry = true
	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.SSA(env)
	return success
}