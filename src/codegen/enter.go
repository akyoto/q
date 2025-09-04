package codegen

import "git.urbach.dev/cli/q/src/asm"

const (
	alignFunction = 0x20
)

// enter sets up the stack frame.
func (f *Function) enter() {
	f.Assembler.Append(&asm.Label{Name: f.FullName, Align: alignFunction})

	if f.Preserved.Count() > 0 && !f.isInit {
		f.Assembler.Append(&asm.Push{Registers: f.Preserved.Slice()})
	}

	if f.hasStackFrame || f.hasExternCalls {
		f.Assembler.Append(&asm.StackFrameStart{FramePointer: f.needsFramePointer, ExternCalls: f.hasExternCalls})
	}
}