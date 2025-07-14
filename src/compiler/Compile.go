package compiler

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/cli/q/src/types"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(build *config.Build) (*core.Environment, error) {
	all, err := scanner.Scan(build)

	if err != nil {
		return nil, err
	}

	// Check for existence of `run.init`
	init := all.Function("run", "init")

	if init == nil {
		return nil, MissingInitFunction
	}

	all.Init = init

	// Check for existence of `main.main`
	main := all.Function("main", "main")

	if main == nil {
		return nil, MissingMainFunction
	}

	all.Main = main

	// Resolve types
	for f := range all.Functions() {
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

	compileFunctions(all.Functions())

	for f := range all.Functions() {
		if f.Err != nil {
			return nil, f.Err
		}
	}

	if build.ShowSSA {
		showSSA(init)
	}

	return all, nil
}