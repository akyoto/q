package compiler

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/errors"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (Result, error) {
	result := Result{}
	all, err := scan(b)

	if err != nil {
		return result, err
	}

	if len(all.Files) == 0 {
		return result, errors.NoInputFiles
	}

	// Resolve the types
	for _, function := range all.Functions {
		err := function.ResolveTypes()

		if err != nil {
			return result, err
		}
	}

	// Parallel compilation
	compileFunctions(all.Functions)

	return result, nil
}