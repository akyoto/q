package build

import (
	"sync"
	"sync/atomic"
)

// Environment represents the global state.
type Environment struct {
	Functions map[string]*Function
}

// NewEnvironment creates a new build environment.
func NewEnvironment() *Environment {
	return &Environment{
		Functions: map[string]*Function{},
	}
}

// ImportDirectory imports a directory to the environment.
func (env *Environment) ImportDirectory(directory string) error {
	files, fileSystemErrors := FindSourceFiles(directory)
	functions, imports, tokenizeErrors := FindFunctions(files)

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

		case directory, ok := <-imports:
			if !ok {
				continue
			}

			err := env.ImportDirectory(directory)

			if err != nil {
				return err
			}

		case function, ok := <-functions:
			if !ok {
				return nil
			}

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
				defer func() {
					if atomic.AddInt64(&function.File.functionCount, -1) == 0 {
						function.File.Close()
					}

					wg.Done()
				}()

				err := Compile(function, env, optimize, verbose)

				if err != nil {
					errors <- err
					return
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
