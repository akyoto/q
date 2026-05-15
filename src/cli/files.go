package cli

import (
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/cli/q/src/verbose"
)

// files shows the entire list of files that are used in a build.
func files(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	env, err := scanner.Scan(b)

	if err != nil {
		return exit(err)
	}

	verbose.Files(env.Files)
	return success
}