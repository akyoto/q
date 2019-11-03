package main

import (
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

// Compiler reads source files and generates machines code.
type Compiler struct {
	assembler *asm.Assembler
}

// NewCompiler creates a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		assembler: asm.New(),
	}
}

// Compile reads the input file and generates an executable binary.
func (compiler *Compiler) Compile(inputFile string, outputFile string) error {
	file := NewFile(inputFile, compiler.assembler)
	err := file.Compile()

	if err != nil {
		return err
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
