package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/set"
)

// Function contains the state that is required to produce assembly code.
type Function struct {
	IR
	CPU               *cpu.CPU
	build             *config.Build
	FullName          string
	Assembler         asm.Assembler
	Preserved         set.Ordered[cpu.Register]
	Count             count
	isInit            bool
	isExit            bool
	needsFramePointer bool
	hasStackFrame     bool
	hasExternCalls    bool
}