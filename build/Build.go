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
	start = time.Now()
	err := build.Environment.ImportDirectory(build.Path, "")

	if err != nil {
		return err
	}

	scan = time.Since(start)

	// Compile
	start = time.Now()
	code, err := build.Compile()

	if err != nil || !build.WriteExecutable {
		return err
	}

	compile = time.Since(start)

	// Write
	start = time.Now()
	err = writeToDisk(code, build.ExecutablePath)

	if err != nil {
		return err
	}

	write = time.Since(start)

	if build.ShowTimings {
		key := log.Faint.Sprint
		log.Info.Printf(key("%-17s")+" %10v μs\n", "Scan files:", scan.Microseconds())
		log.Info.Printf(key("%-17s")+" %10v μs\n", "Compile:", compile.Microseconds())
		log.Info.Printf(key("%-17s")+" %10v μs\n", "Write to disk:", write.Microseconds())
		log.Info.Println(key(strings.Repeat("-", 28)))
		log.Info.Printf(key("%-17s")+color.GreenString(" %10v μs")+"\n", "Total:", (scan + compile + write).Microseconds())
	}

	return nil
}

// Compile compiles all the functions in the environment.
func (build *Build) Compile() (*asm.Assembler, error) {
	_, exists := build.Environment.Functions["main"]

	if !exists {
		return nil, errors.New("Function 'main' has not been defined")
	}

	build.Environment.Compile(build.Optimize, build.ShowAssembly)

	// Generate machine code
	finalCode := asm.New()
	finalCode.Call("main")
	finalCode.Exit(0)

	if !build.WriteExecutable {
		return nil, nil
	}

	for _, function := range build.Environment.Functions {
		if function.Error != nil {
			return nil, function.Error
		}

		if function.File.Error != nil {
			return nil, function.File.Error
		}

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
			log.Info.SetPrefix(log.Faint.Sprint(function.Name) + " ")
			function.assembler.WriteTo(log.Info)
			log.Info.SetPrefix("")
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
