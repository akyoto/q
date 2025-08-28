package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeJump(step *Step, instr *ssa.Jump) {
	f.insertPhiMoves(step)
	f.Assembler.Append(&asm.Jump{Label: instr.To.Label})
}