package build

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/akyoto/directory"
	"github.com/akyoto/q/build/types"
)

// FindFunctions scans the directory for functions.
func FindFunctions(dir string, env *Environment) (<-chan *Function, <-chan *types.Type, <-chan *Import, <-chan error) {
	functions := make(chan *Function, 16)
	structs := make(chan *types.Type, 16)
	imports := make(chan *Import)
	errors := make(chan error)

	go func() {
		findFunctions(dir, env, functions, structs, imports, errors)

		close(functions)
		close(imports)
		close(errors)
	}()

	return functions, structs, imports, errors
}

// findFunctions scans the directory for functions without channel allocations.
func findFunctions(dir string, env *Environment, functions chan<- *Function, structs chan<- *types.Type, imports chan<- *Import, errors chan<- error) {
	wg := sync.WaitGroup{}

	directory.Walk(dir, func(name string) {
		if !strings.HasSuffix(name, ".q") {
			return
		}

		fullPath := filepath.Join(dir, name)
		wg.Add(1)

		go func() {
			defer wg.Done()
			findFunctionsInFile(fullPath, env, functions, structs, imports, errors)
		}()
	})

	wg.Wait()
}

// FindFunctionsInFile a single file for functions.
func FindFunctionsInFile(fileName string, env *Environment) (<-chan *Function, <-chan *types.Type, <-chan *Import, <-chan error) {
	functions := make(chan *Function, 16)
	structs := make(chan *types.Type, 16)
	imports := make(chan *Import)
	errors := make(chan error)

	go func() {
		findFunctionsInFile(fileName, env, functions, structs, imports, errors)

		close(functions)
		close(imports)
		close(errors)
	}()

	return functions, structs, imports, errors
}

// findFunctionsInFile scans the file for functions without channel allocations.
func findFunctionsInFile(fileName string, env *Environment, functions chan<- *Function, structs chan<- *types.Type, imports chan<- *Import, errors chan<- error) {
	file := NewFile(fileName)
	file.environment = env
	err := file.Tokenize()

	if err != nil {
		errors <- err
		return
	}

	err = file.Scan(imports, structs, functions)

	if err != nil {
		errors <- err
		return
	}
}
