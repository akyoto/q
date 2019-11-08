package build

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

// Build describes a compiler build.
type Build struct {
	Path            string
	ExecutablePath  string
	ExecutableName  string
	WriteExecutable bool
	Verbose         bool
	environment
	assembler *asm.Assembler
}

// New creates a new build.
func New(directory string) (*Build, error) {
	directory, err := filepath.Abs(directory)

	if err != nil {
		return nil, err
	}

	executableName := filepath.Base(directory)

	build := &Build{
		Path:            directory,
		ExecutableName:  executableName,
		ExecutablePath:  filepath.Join(directory, executableName),
		WriteExecutable: true,
		assembler:       asm.New(),
		environment: environment{
			functions: map[string]*Function{},
		},
	}

	return build, nil
}

// Run parses the input files and generates an executable binary.
func (build *Build) Run() error {
	files := FindSourceFiles(build.Path)
	functions := FindFunctions(files)

	for function := range functions {
		build.functions[function.Name] = function
	}

	build.Compile()

	if !build.WriteExecutable {
		return nil
	}

	return build.writeToDisk()
}

// writeToDisk writes the executable file to disk.
func (build *Build) writeToDisk() error {
	_, exists := build.functions["main"]

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	// Generate machine code
	build.assembler.Call("main")
	build.assembler.Exit(0)

	// Merge function codes into the main executable
	for _, function := range build.functions {
		build.assembler.Merge(function.compiler.assembler)
	}

	// Produce ELF binary
	binary := elf.New(build.assembler)
	err := binary.WriteToFile(build.ExecutablePath)

	if err != nil {
		return err
	}

	return os.Chmod(build.ExecutablePath, 0755)
}
