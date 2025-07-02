package ssa2asm

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
)

type Compiler struct {
	UniqueName string
	Assembler  asm.Assembler
	CPU        *cpu.CPU
	Count      Count
}