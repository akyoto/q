package compiler

import (
	"maps"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/scanner"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (*core.Environment, error) {
	all, err := scanner.Scan(b)

	if err != nil {
		return nil, err
	}

	if len(all.Files) == 0 {
		return nil, NoInputFiles
	}

	compileFunctions(maps.Values(all.Functions))

	for _, f := range all.Functions {
		if f.Err != nil {
			return nil, f.Err
		}
	}

	if b.ShowSSA {
		showSSA(maps.Values(all.Functions))
	}

	return all, nil
}