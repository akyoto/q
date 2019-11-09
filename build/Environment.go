package build

import (
	"sync"

	"github.com/akyoto/asm"
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
func (env *Environment) Compile() (<-chan *asm.Assembler, <-chan error) {
	assemblers := make(chan *asm.Assembler)
	errors := make(chan error)

	go func() {
		wg := sync.WaitGroup{}

		for _, function := range env.Functions {
			function := function
			wg.Add(1)

			go func() {
				defer wg.Done()
				assembler, err := Compile(function, env)

				if err != nil {
					errors <- err
					return
				}

				if function.TimesUsed == 0 && function.Name != "main" {
					return
				}

				assemblers <- assembler
			}()
		}

		wg.Wait()
		close(assemblers)
		close(errors)
	}()

	return assemblers, errors
}
