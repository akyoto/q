package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Function) GenerateAssembly(ir ssa.IR, stackFrame bool) {
	f.Assembler.Append(&asm.Label{Name: f.FullName})

	if stackFrame {
		f.Assembler.Append(&asm.StackFrameStart{})
	}

	f.Steps = f.createSteps(ir)

	for _, step := range f.Steps {
		f.exec(step)
	}

	if stackFrame {
		f.Assembler.Append(&asm.StackFrameEnd{})
	}

	if f.FullName != "os.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}