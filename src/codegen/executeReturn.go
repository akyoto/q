package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeReturn(instr *ssa.Return) {
	defer f.leave()

	if len(instr.Arguments) == 0 {
		return
	}

	start := len(f.Assembler.Instructions)
	f.moveValuesToRegisters(instr.Arguments, f.CPU.Call.Out)
	reorderMoves(f.Assembler.Instructions[start:])
}