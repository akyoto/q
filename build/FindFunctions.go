package build

import (
	"sync"
)

// FindFunctions scans the files for functions.
func FindFunctions(files <-chan *File) (<-chan *Function, <-chan error) {
	functions := make(chan *Function)
	errors := make(chan error)

	go func() {
		wg := sync.WaitGroup{}

		for file := range files {
			file := file
			wg.Add(1)

			go func() {
				defer wg.Done()
				err := file.Tokenize()

				if err != nil {
					errors <- err
					return
				}

				err = file.Scan(functions)

				if err != nil {
					errors <- err
					return
				}
			}()
		}

		wg.Wait()
		close(functions)
		close(errors)
	}()

	return functions, errors
}
