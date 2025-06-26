package compiler

import (
	"maps"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/cli/q/src/types"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(b *build.Build) (*core.Environment, error) {
	all, err := scanner.Scan(b)

	if err != nil {
		return nil, err
	}

	// Check for existence of `init`
	init, exists := all.Functions["core.init"]

	if !exists {
		return nil, MissingInitFunction
	}

	// Check for existence of `main`
	_, exists = all.Functions["main.main"]

	if !exists {
		return nil, MissingMainFunction
	}

	// Resolve types
	for _, f := range all.Functions {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Typ = types.Parse(input.Source[1:], f.File.Bytes)
			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			if len(output.Source) > 1 {
				output.Typ = types.Parse(output.Source[1:], f.File.Bytes)
			} else {
				output.Typ = types.Parse(output.Source, f.File.Bytes)
			}

			f.Type.Output[i] = output.Typ
		}
	}

	compileFunctions(maps.Values(all.Functions))

	for _, f := range all.Functions {
		if f.Err != nil {
			return nil, f.Err
		}
	}

	if b.ShowSSA {
		showSSA(init)
	}

	return all, nil
}