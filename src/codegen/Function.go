package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
)

type Function struct {
	FullName          string
	Assembler         asm.Assembler
	Steps             []*step
	ValueToStep       map[ssa.Value]*step
	CPU               *cpu.CPU
	Preserved         set.Ordered[cpu.Register]
	Count             count
	isInit            bool
	isExit            bool
	needsFramePointer bool
	hasStackFrame     bool
	hasExternCalls    bool
}