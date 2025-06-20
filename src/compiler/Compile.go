package compiler

import (
	"maps"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/scanner"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (Result, error) {
	result := Result{}
	all, err := scanner.Scan(b)

	if err != nil {
		return result, err
	}

	if len(all.Files) == 0 {
		return result, NoInputFiles
	}

	for _, function := range all.Functions {
		err := function.ResolveTypes()

		if err != nil {
			return result, err
		}
	}

	compileFunctions(maps.Values(all.Functions))
	return result, nil
}