package build

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
	"github.com/akyoto/q/build/log"
)

// Build describes a compiler build.
type Build struct {
	Path            string
	ExecutablePath  string
	ExecutableName  string
	Environment     *Environment
	WriteExecutable bool
	Optimize        bool
	Verbose         bool
}

// New creates a new build.
func New(directory string) (*Build, error) {
	directory, err := filepath.Abs(directory)

	if err != nil {
		return nil, err
	}

	executableName := filepath.Base(directory)
	environment, err := NewEnvironment()

	if err != nil {
		return nil, err
	}

	build := &Build{
		Path:            directory,
		ExecutableName:  executableName,
		ExecutablePath:  filepath.Join(directory, executableName),
		WriteExecutable: true,
		Environment:     environment,
	}

	return build, nil
}

// Run parses the input files and generates an executable binary.
func (build *Build) Run() error {
	err := build.Environment.ImportDirectory(build.Path, "")

	if err != nil {
		return err
	}

	return build.Compile()
}

// Compile compiles all the functions in the environment.
func (build *Build) Compile() error {
	_, exists := build.Environment.Functions["main"]

	if !exists {
		return errors.New("Function 'main' has not been defined")
	}

	var results []*Function
	resultsChannel, errors := build.Environment.Compile(build.Optimize, build.Verbose)

	// Generate machine code
	finalCode := asm.New()
	finalCode.Call("main")
	finalCode.Exit(0)

	for {
		select {
		case err, ok := <-errors:
			if ok {
				return err
			}

		case compiled, ok := <-resultsChannel:
			if !ok {
				goto done
			}

			results = append(results, compiled)
		}
	}

done:
	if !build.WriteExecutable {
		return nil
	}

	for _, function := range results {
		if function.CallCount == 0 {
			continue
		}

		if function.Name != "main" && function.CanInline() {
			continue
		}

		// Merge function code into the main finalCode
		finalCode.Merge(function.assembler.Finalize())

		// Show assembler code of used functions
		if build.Verbose {
			function.assembler.WriteTo(log.Info)
			log.Info.Println()
		}
	}

	for _, err := range finalCode.Verify() {
		return err
	}

	return writeToDisk(finalCode, build.ExecutablePath)
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
