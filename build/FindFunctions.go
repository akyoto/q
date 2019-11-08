package build

import (
	"os"
	"sync"
	"sync/atomic"
)

// FindFunctions scans the files for functions.
func FindFunctions(files <-chan *File) <-chan *Function {
	functions := make(chan *Function)

	go func() {
		wg := sync.WaitGroup{}
		errorCount := uint64(0)

		for file := range files {
			file := file
			wg.Add(1)

			go func() {
				defer wg.Done()
				err := file.Tokenize()

				if err != nil {
					stderr.Println(err)
					atomic.AddUint64(&errorCount, 1)
					return
				}

				err = file.Scan(functions)

				if err != nil {
					stderr.Println(err)
					atomic.AddUint64(&errorCount, 1)
					return
				}
			}()
		}

		wg.Wait()

		if errorCount > 0 {
			os.Exit(1)
		}

		close(functions)
	}()

	return functions
}
