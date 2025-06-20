package compiler

import (
	"iter"
	"sync"

	"git.urbach.dev/cli/q/src/core"
)

// compileFunctions starts a goroutine for each function compilation and waits for completion.
func compileFunctions(functions iter.Seq[*core.Function]) {
	wg := sync.WaitGroup{}

	for function := range functions {
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