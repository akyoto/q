package build

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/log"
)

// Environment represents the global state.
type Environment struct {
	functions map[string]*Function
}

// NewEnvironment creates a new build environment.
func NewEnvironment() *Environment {
	return &Environment{
		functions: map[string]*Function{},
	}
}

// Compile compiles all functions.
func (env *Environment) Compile() <-chan *asm.Assembler {
	assemblers := make(chan *asm.Assembler)

	go func() {
		wg := sync.WaitGroup{}
		errorCount := uint64(0)

		for _, function := range env.functions {
			function := function
			wg.Add(1)

			go func() {
				defer wg.Done()
				assembler, err := Compile(function, env)

				if err != nil {
					log.Error.Println(err)
					atomic.AddUint64(&errorCount, 1)
					return
				}

				assemblers <- assembler
			}()
		}

		wg.Wait()

		if errorCount > 0 {
			os.Exit(1)
		}

		close(assemblers)
	}()

	return assemblers
}
