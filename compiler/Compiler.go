package compiler

import (
	"errors"
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
		assembler:       assemblerPool.Get().(*asm.Assembler),
		WriteExecutable: true,
	}
}

// Compile reads the input file and generates an executable binary.
func (compiler *Compiler) Compile(inputFile string, outputFile string) error {
	// Call main and exit
	compiler.assembler.Call("main")
	compiler.assembler.Exit(0)

	file := NewFile(inputFile, compiler.assembler)
	defer file.Close()
	err := file.Compile()

	if err != nil {
		return err
	}

	if !compiler.WriteExecutable {
		return nil
	}

	if file.functions["main"] == nil {
		return errors.New("Function 'main' has not been defined")
	}

	// Produce ELF binary
	binary := elf.New(compiler.assembler)
	err = binary.WriteToFile(outputFile)

	if err != nil {
		return err
	}

	return os.Chmod(outputFile, 0755)
}

// Close frees up resources used by the compiler.
func (compiler *Compiler) Close() {
	compiler.assembler.Reset()
	assemblerPool.Put(compiler.assembler)
}
