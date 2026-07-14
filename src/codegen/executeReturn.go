package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeReturn(instr *ssa.Return) {
	defer f.leave()

	if len(instr.Arguments) == 0 {
		return
	}

	f.moveValuesToRegisters(instr.Arguments, f.CPU.Call.Out)
}