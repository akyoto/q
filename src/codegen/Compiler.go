package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

type Compiler struct {
	FullName    string
	Assembler   asm.Assembler
	Steps       []*Step
	ValueToStep map[ssa.Value]*Step
	CPU         *cpu.CPU
	Count       Count
}