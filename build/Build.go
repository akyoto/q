package build

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

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
		assembler:       assemblerPool.Get().(*asm.Assembler),
	}

	build.ExecutablePath = filepath.Join(build.Path, build.ExecutableName)
	return build, nil
}

// Run parses the input files and generates an executable binary.
func (build *Build) Run() error {
	files := []*File{}

	err := filepath.Walk(build.Path, func(path string, info os.FileInfo, err error) error {
		if path == build.Path {
			return nil
		}

		if info.IsDir() {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".q") {
			return nil
		}

		if build.Verbose {
			fmt.Println("Scanning", info.Name())
		}

		file := NewFile(path)
		files = append(files, file)
		err = file.Read()

		if err != nil {
			return err
		}

		return file.Scan(func(function *Function) {
			function.compiler.environment = &build.Environment
			build.functions.Store(function.Name, function)
		})
	})

	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("No source files found in %s", build.Path)
	}

	if !build.WriteExecutable {
		return nil
	}

	_, exists := build.functions.Load("main")

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	// Generate machine code
	build.assembler.Call("main")
	build.assembler.Exit(0)

	// Start parallel function compilation
	wg := sync.WaitGroup{}

	build.functions.Range(func(key interface{}, value interface{}) bool {
		function := value.(*Function)
		wg.Add(1)

		go func() {
			err := function.Compile()

			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(1)
			}

			wg.Done()
		}()

		return true
	})

	wg.Wait()

	// Merge function codes into the main executable
	build.functions.Range(func(key interface{}, value interface{}) bool {
		function := value.(*Function)
		build.assembler.Merge(function.compiler.assembler)
		return true
	})

	// Close files
	for _, file := range files {
		file.Close()
	}

	// Produce ELF binary
	binary := elf.New(build.assembler)
	err = binary.WriteToFile(build.ExecutablePath)

	if err != nil {
		return err
	}

	return os.Chmod(build.ExecutablePath, 0755)
}

// Close frees up resources used by the build.
func (build *Build) Close() {
	build.assembler.Reset()
	assemblerPool.Put(build.assembler)

	build.functions.Range(func(key interface{}, value interface{}) bool {
		assembler := value.(*Function).compiler.assembler
		assembler.Reset()
		assemblerPool.Put(assembler)
		return true
	})
}
