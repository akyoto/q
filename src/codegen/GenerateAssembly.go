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

	f.createSteps(ir)

	if f.Preserved.Count() > 0 {
		f.Assembler.Append(&asm.PushRegisters{Registers: f.Preserved.Slice()})
	}

	for _, step := range f.Steps {
		f.exec(step)
	}

	if f.FullName == "os.exit" {
		return
	}

	if f.Preserved.Count() > 0 {
		f.Assembler.Append(&asm.PopRegisters{Registers: f.Preserved.Slice()})
	}

	if stackFrame {
		f.Assembler.Append(&asm.StackFrameEnd{})
	}

	f.Assembler.Append(&asm.Return{})
}