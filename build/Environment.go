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

// Compile compiles all functions.
func (env *Environment) Compile(optimize bool, verbose bool) (<-chan *CompilationResult, <-chan error) {
	results := make(chan *CompilationResult)
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

				assembler, err := Compile(function, env, optimize, verbose)

				if err != nil {
					errors <- err
					return
				}

				results <- &CompilationResult{function, assembler}
			}()
		}

		wg.Wait()
		close(results)
		close(errors)
	}()

	return results, errors
}
