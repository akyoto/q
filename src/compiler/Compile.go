package compiler

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/cli/q/src/verbose"
)

// Compile waits for the scan to finish and compiles all functions.
func Compile(build *config.Build) (*core.Environment, error) {
	env, err := scanner.Scan(build)

	if err != nil {
		return nil, err
	}

	// Check for existence of `run.init`
	init := env.Function("run", "init")

	if init == nil {
		return nil, MissingInitFunction
	}

	env.Init = init

	// Check for existence of `main.main`
	main := env.Function("main", "main")

	if main == nil {
		return nil, MissingMainFunction
	}

	env.Main = main

	// Resolve types
	for f := range env.Functions() {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Typ = types.Parse(input.Tokens[1:], f.File.Bytes)
			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			if len(output.Tokens) > 1 {
				output.Typ = types.Parse(output.Tokens[1:], f.File.Bytes)
			} else {
				output.Typ = types.Parse(output.Tokens, f.File.Bytes)
			}

			f.Type.Output[i] = output.Typ
		}
	}

	// Start parallel compilation
	compileFunctions(env.Functions())

	// Report errors if any occurred
	for f := range env.Functions() {
		if f.Err != nil {
			return nil, f.Err
		}
	}

	// Check for unused imports in all files
	for _, file := range env.Files {
		for _, imp := range file.Imports {
			if imp.Used.Load() == 0 {
				return nil, errors.New(&UnusedImport{Package: imp.Package}, file, imp.Position)
			}
		}
	}

	// Verbose output
	if build.ShowIR {
		if build.ShowHeaders {
			verbose.Header(verbose.HeaderIR)
		}

		verbose.IR(env.Init)
	}

	if build.ShowASM {
		if build.ShowHeaders {
			verbose.Header(verbose.HeaderASM)
		}

		verbose.ASM(env.Init)
	}

	return env, nil
}