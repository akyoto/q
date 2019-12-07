package build

import (
	"sync"
)

// FindFunctions scans the files for functions.
func FindFunctions(files <-chan *File, stdLib string) (<-chan *Function, <-chan *Import, <-chan error) {
	functions := make(chan *Function, 16)
	imports := make(chan *Import)
	errors := make(chan error)

	go func() {
		wg := sync.WaitGroup{}

		for file := range files {
			file := file
			file.standardLibrary = stdLib
			wg.Add(1)

			go func() {
				defer wg.Done()
				err := file.Tokenize()

				if err != nil {
					errors <- err
					return
				}

				err = file.Scan(imports, functions)

				if err != nil {
					errors <- err
					return
				}
			}()
		}

		wg.Wait()
		close(functions)
		close(imports)
		close(errors)
	}()

	return functions, imports, errors
}
