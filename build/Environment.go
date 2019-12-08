package build

import (
	"os"
	"sync"
	"sync/atomic"
)

// Environment represents the global state.
type Environment struct {
	Packages        map[string]bool
	Functions       map[string]*Function
	Types           map[string]*Type
	StandardLibrary string
}

// NewEnvironment creates a new build environment.
func NewEnvironment() (*Environment, error) {
	stdLib, err := FindStandardLibrary()

	if err != nil {
		return nil, err
	}

	environment := &Environment{
		Packages:  map[string]bool{},
		Functions: map[string]*Function{},
		Types: map[string]*Type{
			"int64": {Name: "int64", Size: 8},
			"int32": {Name: "int32", Size: 4},
			"int16": {Name: "int16", Size: 2},
			"int8":  {Name: "int8", Size: 1},
		},
		StandardLibrary: stdLib,
	}

	return environment, nil
}

// ImportDirectory imports a directory to the environment.
func (env *Environment) ImportDirectory(directory string, prefix string) error {
	_, err := os.Stat(directory)

	if err != nil {
		return err
	}

	files, fileSystemErrors := FindSourceFiles(directory)
	functions, imports, tokenizeErrors := FindFunctions(files, env.StandardLibrary)

	for {
		select {
		case err, ok := <-fileSystemErrors:
			if ok {
				return err
			}

		case err, ok := <-tokenizeErrors:
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
