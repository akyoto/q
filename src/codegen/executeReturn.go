package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeReturn(instr *ssa.Return) {
	defer f.leave()

	if len(instr.Arguments) == 0 {
		return
	}

	for i, arg := range instr.Arguments {
		retVal := f.ValueToStep[arg]

		if retVal.Register == -1 || retVal.Register == f.CPU.Call.Out[i] {
			return
		}

		f.Assembler.Append(&asm.Move{
			Destination: f.CPU.Call.Out[i],
			Source:      retVal.Register,
		})
	}
}