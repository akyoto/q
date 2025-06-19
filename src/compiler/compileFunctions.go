package compiler

import (
	"sync"

	"git.urbach.dev/cli/q/src/core"
)

// compileFunctions starts a goroutine for each function compilation and waits for completion.
func compileFunctions(functions map[string]*core.Function) {
	wg := sync.WaitGroup{}

	for _, function := range functions {
		if function.IsExtern() {
			continue
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			function.Compile()
		}()
	}

	wg.Wait()
}