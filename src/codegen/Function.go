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
	ValueToStep       map[ssa.Value]*Step
	BlockToRegion     map[*ssa.Block]region
	CPU               *cpu.CPU
	build             *config.Build
	FullName          string
	Assembler         asm.Assembler
	Steps             []*Step
	Preserved         set.Ordered[cpu.Register]
	Count             count
	isInit            bool
	isExit            bool
	needsFramePointer bool
	hasStackFrame     bool
	hasExternCalls    bool
}