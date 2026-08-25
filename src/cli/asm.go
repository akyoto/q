package cli

import (
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/verbose"
)

// asm shows the assembly code.
func asm(args []string) int {
	pattern, args, err := filter(args)

	if err != nil {
		return exit(err)
	}

	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	b.Dry = true
	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.ASM(env, pattern)
	return success
}