package compiler

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/scanner"
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

	// Parse struct field types and calculate the size of all structs.
	// We couldn't do that during the scan phase because it's possible
	// that a field references a type that will only be known after the
	// full scan is finished.
	err = parseFieldTypes(env.Structs(), env)

	if err != nil {
		return nil, err
	}

	// Parse input and output types so we have type information
	// ready for all functions before parallel compilation starts.
	// This ensures that the function compilers have access to
	// type checking for all function calls.
	err = parseTypes(env.Functions(), env)

	if err != nil {
		return nil, err
	}

	// Start parallel compilation of all functions.
	// We compile every function for syntax checks even if
	// they are thrown away later during dead code elimination.
	parallel(env.Functions(), func(f *core.Function) { f.Compile() })

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

	// Now that we know which functions are alive, start parallel
	// assembly code generation only for the live functions.
	parallel(env.LiveFunctions(), func(f *core.Function) { f.Assemble() })

	return env, nil
}