package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/set"
)

// Function contains the state that is required to produce assembly code.
type Function struct {
	IR
	CPU               *cpu.CPU
	arch              arch
	FullName          string
	Assembler         asm.Assembler
	Preserved         set.Ordered[cpu.Register]
	Count             count
	stackSize         uint
	IsExit            bool
	needsFramePointer bool
	hasStackFrame     bool
	hasExternCalls    bool
}