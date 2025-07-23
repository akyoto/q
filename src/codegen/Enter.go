package codegen

import "git.urbach.dev/cli/q/src/asm"

// Enter sets up the stack frame.
func (f *Function) Enter() {
	f.Assembler.Append(&asm.Label{Name: f.FullName})

	if f.Preserved.Count() > 0 && !f.isInit {
		f.Assembler.Append(&asm.Push{Registers: f.Preserved.Slice()})
	}

	if f.hasStackFrame || f.hasExternCalls {
		f.Assembler.Append(&asm.StackFrameStart{FramePointer: f.needsFramePointer, ExternCalls: f.hasExternCalls})
	}
}