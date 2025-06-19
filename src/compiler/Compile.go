package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/scanner"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (Result, error) {
	result := Result{}
	files, errors := scanner.Scan(b)

	for files != nil || errors != nil {
		select {
		case file, ok := <-files:
			if !ok {
				files = nil
				continue
			}

			fmt.Println(file.Path)

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}

			return result, err
		}
	}

	return result, nil
}