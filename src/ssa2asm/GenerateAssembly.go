package ssa2asm

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Compiler) GenerateAssembly(ir ssa.IR, stackFrame bool) {
	f.Assembler.Append(&asm.Label{Name: f.UniqueName})

	if stackFrame {
		f.Assembler.Append(&asm.StackFrameStart{})
	}

	f.Steps = f.CreateSteps(ir)

	for _, step := range f.Steps {
		f.Exec(step)
	}

	if stackFrame {
		f.Assembler.Append(&asm.StackFrameEnd{})
	}

	if f.UniqueName != "os.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}