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
	Environment     *Environment
	WriteExecutable bool
	Verbose         bool
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
	files, fileSystemErrors := FindSourceFiles(build.Path)
	functions, tokenizeErrors := FindFunctions(files)

	for {
		select {
		case err, ok := <-fileSystemErrors:
			if ok {
				return err
			}

		case err, ok := <-tokenizeErrors:
			if ok {
				return err
			}

		case function, ok := <-functions:
			if !ok {
				goto done
			}

			build.Environment.Functions[function.Name] = function
		}
	}

done:
	return build.Compile()
}

// Compile compiles all the functions in the environment.
func (build *Build) Compile() error {
	_, exists := build.Environment.Functions["main"]

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	assemblers, errors := build.Environment.Compile()

	// Generate machine code
	main := asm.New()
	main.Call("main")
	main.Exit(0)

	for {
		select {
		case err, ok := <-errors:
			if ok {
				return err
			}

		case functionCode, ok := <-assemblers:
			if !ok {
				goto done
			}

			// Merge function code into the main executable
			main.Merge(functionCode)
		}
	}

done:
	if !build.WriteExecutable {
		return nil
	}

	for _, err := range main.Verify() {
		return err
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
