package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
)

// Function contains the state that is required to produce assembly code.
type Function struct {
	FullName          string
	Assembler         asm.Assembler
	Steps             []*Step
	ValueToStep       map[ssa.Value]*Step
	CPU               *cpu.CPU
	Preserved         set.Ordered[cpu.Register]
	Count             count
	build             *config.Build
	isInit            bool
	isExit            bool
	needsFramePointer bool
	hasStackFrame     bool
	hasExternCalls    bool
}