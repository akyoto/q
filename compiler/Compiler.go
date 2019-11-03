package compiler

import (
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

// Compiler reads source files and generates machines code.
type Compiler struct {
	assembler       *asm.Assembler
	WriteExecutable bool
}

// New creates a new compiler.
func New() *Compiler {
	return &Compiler{
		assembler:       asm.New(),
		WriteExecutable: true,
	}
}

// Compile reads the input file and generates an executable binary.
func (compiler *Compiler) Compile(inputFile string, outputFile string) error {
	file := NewFile(inputFile, compiler.assembler)
	err := file.Compile()

	if err != nil {
		return err
	}

	if !compiler.WriteExecutable {
		return nil
	}

	// Programs should always exit
	compiler.assembler.Exit(0)

	// Produce ELF binary
	binary := elf.New(compiler.assembler)
	err = binary.WriteToFile(outputFile)

	if err != nil {
		return err
	}

	return os.Chmod(outputFile, 0755)
}
