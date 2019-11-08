package build

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/akyoto/q/build/log"
)

// environment represents the global state.
type environment struct {
	functions map[string]*Function
}

// Compile compiles all functions.
func (env *environment) Compile() {
	wg := sync.WaitGroup{}
	errorCount := uint64(0)

	for _, function := range env.functions {
		function := function
		function.compiler.environment = env
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := function.Compile()

			if err != nil {
				log.Error.Println(err)
				atomic.AddUint64(&errorCount, 1)
			}
		}()
	}

	wg.Wait()

	if errorCount > 0 {
		os.Exit(1)
	}
}
