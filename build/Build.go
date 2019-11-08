package build

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

// Build describes a compiler build.
type Build struct {
	Environment
	Path            string
	ExecutablePath  string
	ExecutableName  string
	WriteExecutable bool
	Verbose         bool
	assembler       *asm.Assembler
}

// New creates a new build.
func New(directory string) (*Build, error) {
	directory, err := filepath.Abs(directory)

	if err != nil {
		return nil, err
	}

	build := &Build{
		Path:            directory,
		ExecutableName:  filepath.Base(directory),
		WriteExecutable: true,
		assembler:       asm.New(),
	}

	build.ExecutablePath = filepath.Join(build.Path, build.ExecutableName)
	return build, nil
}

// Run parses the input files and generates an executable binary.
func (build *Build) Run() error {
	files := FindSourceFiles(build.Path)
	functions := FindFunctions(files)

	for function := range functions {
		function.compiler.environment = &build.Environment
		build.functions.Store(function.Name, function)
	}

	if !build.WriteExecutable {
		return nil
	}

	return build.WriteToDisk()
}

// WriteToDisk writes the executable file to disk.
func (build *Build) WriteToDisk() error {
	_, exists := build.functions.Load("main")

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	// Generate machine code
	build.assembler.Call("main")
	build.assembler.Exit(0)

	// Start parallel function compilation
	wg := sync.WaitGroup{}
	errorCount := uint64(0)

	build.functions.Range(func(key interface{}, value interface{}) bool {
		function := value.(*Function)
		wg.Add(1)

		go func() {
			err := function.Compile()

			if err != nil {
				stderr.Println(err)
				atomic.AddUint64(&errorCount, 1)
			}

			wg.Done()
		}()

		return true
	})

	wg.Wait()

	if errorCount > 0 {
		os.Exit(1)
	}

	// Merge function codes into the main executable
	build.functions.Range(func(key interface{}, value interface{}) bool {
		function := value.(*Function)
		build.assembler.Merge(function.compiler.assembler)
		return true
	})

	// Produce ELF binary
	binary := elf.New(build.assembler)
	err := binary.WriteToFile(build.ExecutablePath)

	if err != nil {
		return err
	}

	return os.Chmod(build.ExecutablePath, 0755)
}
