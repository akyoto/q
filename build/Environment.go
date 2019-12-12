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
	functions, imports, errors := FindFunctions(directory, env)
	return env.Import(prefix, functions, imports, errors)
}

// Import imports the given functions and imports to the environment.
func (env *Environment) Import(prefix string, functions <-chan *Function, imports <-chan *Import, errors <-chan error) error {
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
func (env *Environment) Compile(optimize bool, verbose bool) (<-chan *Function, <-chan error) {
	results := make(chan *Function)
	errors := make(chan error)

	go func() {
		wg := sync.WaitGroup{}

		for _, function := range env.Functions {
			function := function
			wg.Add(1)

			go func() {
				defer wg.Done()
				err := Compile(function, env, optimize, verbose)

				if err != nil {
					errors <- err
					return
				}

				if atomic.AddInt64(&function.File.functionCount, -1) == 0 {
					err := function.File.Close()

					if err != nil {
						errors <- err
						return
					}
				}

				results <- function
			}()
		}

		wg.Wait()
		close(results)
		close(errors)
	}()

	return results, errors
}
