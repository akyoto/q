package compiler

import (
	"iter"
	"sync"

	"git.urbach.dev/cli/q/src/core"
)

// parallel starts a goroutine for each function and waits for completion.
func parallel(functions iter.Seq[*core.Function], call func(*core.Function)) {
	wg := sync.WaitGroup{}

	for function := range functions {
		if function.IsExtern() {
			continue
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			call(function)
		}()
	}

	wg.Wait()
}