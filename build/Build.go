package build

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
	"github.com/akyoto/color"
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
	ShowTimings     bool
	ShowAssembly    bool
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
	var (
		start   time.Time
		scan    time.Duration
		compile time.Duration
		write   time.Duration
	)

	// Scan
	if build.ShowTimings {
		start = time.Now()
	}

	err := build.Environment.ImportDirectory(build.Path, "")

	if err != nil {
		return err
	}

	if build.ShowTimings {
		scan = time.Since(start)
	}

	// Compile
	if build.ShowTimings {
		start = time.Now()
	}

	code, err := build.Compile()

	if err != nil || !build.WriteExecutable {
		return err
	}

	if build.ShowTimings {
		compile = time.Since(start)
	}

	// Write
	if build.ShowTimings {
		start = time.Now()
	}

	err = writeToDisk(code, build.ExecutablePath)

	if err != nil {
		return err
	}

	if build.ShowTimings {
		write = time.Since(start)

		key := color.New(color.Faint).Sprint
		log.Info.Printf(key("%-17s")+" %10v\n", "Scan files:", scan)
		log.Info.Printf(key("%-17s")+" %10v\n", "Compile:", compile)
		log.Info.Printf(key("%-17s")+" %10v\n", "Write to disk:", write)
		log.Info.Println(key(strings.Repeat("-", 28)))
		log.Info.Printf(key("%-17s")+color.GreenString(" %10v")+"\n", "Total:", scan+compile+write)
	}

	return nil
}

// Compile compiles all the functions in the environment.
func (build *Build) Compile() (*asm.Assembler, error) {
	_, exists := build.Environment.Functions["main"]

	if !exists {
		return nil, errors.New("Function 'main' has not been defined")
	}

	var results []*Function
	resultsChannel, errors := build.Environment.Compile(build.Optimize, build.ShowAssembly)

	// Generate machine code
	finalCode := asm.New()
	finalCode.Call("main")
	finalCode.Exit(0)

	for {
		select {
		case err, ok := <-errors:
			if ok {
				return nil, err
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
		return nil, nil
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
		if build.ShowAssembly {
			function.assembler.WriteTo(log.Info)
			log.Info.Println()
		}
	}

	for _, err := range finalCode.Verify() {
		return nil, err
	}

	return finalCode, nil
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
