package compiler

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/scanner"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (Result, error) {
	scanner.Scan(b)
	return Result{}, nil
}