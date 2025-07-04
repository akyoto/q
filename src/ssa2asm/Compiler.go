package ssa2asm

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

type Compiler struct {
	UniqueName  string
	Assembler   asm.Assembler
	ValueToStep map[ssa.Value]*Step
	CPU         *cpu.CPU
	Count       Count
}