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
	*Environment
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
		Environment:     NewEnvironment(),
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

	_, exists := build.functions["main"]

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	assemblers := build.Compile()

	// Generate machine code
	main := asm.New()
	main.Call("main")
	main.Exit(0)

	// Merge function codes into the main executable
	for functionCode := range assemblers {
		main.Merge(functionCode)
	}

	if !build.WriteExecutable {
		return nil
	}

	return writeToDisk(main, build.ExecutablePath)
}

// writeToDisk writes the executable file to disk.
func writeToDisk(main *asm.Assembler, filePath string) error {
	binary := elf.New(main)
	err := binary.WriteToFile(filePath)

	if err != nil {
		return err
	}

	return os.Chmod(filePath, 0755)
}
