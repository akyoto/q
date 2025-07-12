package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

type Function struct {
	FullName    string
	Assembler   asm.Assembler
	Steps       []*step
	ValueToStep map[ssa.Value]*step
	CPU         *cpu.CPU
	Count       count
}