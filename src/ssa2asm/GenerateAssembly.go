package ssa2asm

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Compiler) GenerateAssembly(ir ssa.IR, isLeaf bool) {
	f.Assembler.Append(&asm.Label{Name: f.UniqueName})

	if !isLeaf && f.UniqueName != "run.init" {
		f.Assembler.Append(&asm.FunctionStart{})
	}

	f.Steps = f.CreateSteps(ir)

	for _, step := range f.Steps {
		f.Exec(&step)
	}

	if !isLeaf && f.UniqueName != "run.init" {
		f.Assembler.Append(&asm.FunctionEnd{})
	}

	if f.UniqueName != "os.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}