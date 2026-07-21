package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
)

// enter sets up the stack frame.
func (f *Function) enter() {
	f.Assembler.Append(&asm.Label{
		Name:  f.FullName,
		Align: f.alignment(),
	})

	if f.Preserved.Count() > 0 && !f.IsExit {
		f.Assembler.Append(&asm.Push{Registers: f.Preserved.Slice()})
	}

	if f.hasExternCalls || (f.hasStackFrame && !f.IsExit) {
		f.Assembler.Append(&asm.StackFrameStart{FramePointer: f.needsFramePointer, ExternCalls: f.hasExternCalls})
	}

	if f.stackSize > 0 {
		f.stackSize = (f.stackSize + 15) &^ 15

		f.Assembler.Append(&asm.SubtractNumber{
			Destination: f.CPU.StackPointer,
			Source:      f.CPU.StackPointer,
			Number:      int(f.stackSize),
		})
	}
}