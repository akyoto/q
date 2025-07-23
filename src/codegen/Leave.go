package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
)

// Leave cleans up the stack and returns to the caller.
func (f *Function) Leave() {
	if f.isExit {
		return
	}

	if f.hasStackFrame || f.hasExternCalls {
		f.Assembler.Append(&asm.StackFrameEnd{FramePointer: f.needsFramePointer})
	}

	if f.isInit {
		return
	}

	if f.Preserved.Count() > 0 {
		f.Assembler.Append(&asm.Pop{Registers: f.Preserved.Slice()})
	}

	f.Assembler.Append(&asm.Return{})
}