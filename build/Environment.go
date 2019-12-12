package build

import (
	"sync"
	"sync/atomic"

	"github.com/akyoto/q/build/types"
)

// Environment represents the global state.
type Environment struct {
	Packages        map[string]bool
	Functions       map[string]*Function
	Types           map[string]*types.Type
	StandardLibrary string
}

// NewEnvironment creates a new build environment.
func NewEnvironment() (*Environment, error) {
	standardLibrary, err := FindStandardLibrary()

	if err != nil {
		return nil, err
	}

	environment := &Environment{
		Packages:        map[string]bool{},
		Functions:       map[string]*Function{},
		Types:           types.Default,
		StandardLibrary: standardLibrary,
	}

	return environment, nil
}

// ImportDirectory imports a directory to the environment.
func (env *Environment) ImportDirectory(directory string, prefix string) error {
	functions, structs, imports, errors := FindFunctions(directory, env)
	return env.Import(prefix, functions, structs, imports, errors)
}

// Import imports the given functions and imports to the environment.
func (env *Environment) Import(prefix string, functions <-chan *Function, structs <-chan *types.Type, imports <-chan *Import, errors <-chan error) error {
	for {
		select {
		case err, ok := <-errors:
			if ok {
				return err
			}

		case imp, ok := <-imports:
			if !ok {
				continue
			}

			if env.Packages[imp.Path] {
				continue
			}

			env.Packages[imp.Path] = true
			err := env.ImportDirectory(imp.FullPath, imp.Path+".")

			if err != nil {
				return err
			}

		case typ, ok := <-structs:
			if !ok {
				return nil
			}

			typ.Name = prefix + typ.Name
			env.Types[typ.Name] = typ

		case function, ok := <-functions:
			if !ok {
				return nil
			}

			function.Name = prefix + function.Name
			env.Functions[function.Name] = function
		}
	}
}

// Compile compiles all functions.
func (env *Environment) Compile(optimize bool, verbose bool) {
	wg := sync.WaitGroup{}

	for _, function := range env.Functions {
		wg.Add(1)

		go func(function *Function) {
			defer wg.Done()
			Compile(function, env, optimize, verbose)

			if function.Error != nil {
				return
			}

			if atomic.AddInt64(&function.File.functionCount, -1) == 0 {
				function.File.Close()
			}
		}(function)
	}

	wg.Wait()
}
