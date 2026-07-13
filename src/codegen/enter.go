package codegen

import (
	"fmt"
	"os"

	"git.urbach.dev/cli/q/src/asm"
)

// enter sets up the stack frame.
func (f *Function) enter() {
	f.Assembler.Append(&asm.Label{
		Name:  f.FullName,
		Align: f.alignment(),
	})

	if !f.IsExit {
		if f.Preserved.Count() > 0 {
			f.Assembler.Append(&asm.Push{Registers: f.Preserved.Slice()})
		}

		if f.hasStackFrame || f.hasExternCalls {
			f.Assembler.Append(&asm.StackFrameStart{FramePointer: f.needsFramePointer, ExternCalls: f.hasExternCalls})
		}
	}

	if f.stackSize > 0 {
		fmt.Fprintf(os.Stderr, "%s warning: register spills to memory are not fully implemented yet (try to use less live values)\n", f.FullName)
		f.stackSize = (f.stackSize + 15) &^ 15

		f.Assembler.Append(&asm.SubtractNumber{
			Destination: f.CPU.StackPointer,
			Source:      f.CPU.StackPointer,
			Number:      int(f.stackSize),
		})
	}
}