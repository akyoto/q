package compiler

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/errors"
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
		return result, errors.NoInputFiles
	}

	for _, function := range all.Functions {
		err := function.ResolveTypes()

		if err != nil {
			return result, err
		}
	}

	compileFunctions(all.Functions)
	return result, nil
}