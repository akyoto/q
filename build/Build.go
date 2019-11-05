package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	build.assembler.Call("main")
	build.assembler.Exit(0)

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
			fmt.Println("Compiling", info.Name())
		}

		file := NewFile(path, build.assembler)
		compilerError := file.Compile()
		file.Close()
		files = append(files, file)
		return compilerError
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

	// if file.functions["main"] == nil {
	// 	return errors.New("Function 'main' has not been defined")
	// }

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
}
