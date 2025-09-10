package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeLoad(step *Step, instr *ssa.Load) {
	if step.Register == -1 {
		return
	}

	address := f.ValueToStep[instr.Address]
	index := f.ValueToStep[instr.Index]
	elementSize := step.Value.Type().Size()

	f.Assembler.Append(&asm.Load{
		Base:        address.Register,
		Index:       index.Register,
		Destination: step.Register,
		Scale:       instr.Scale,
		Length:      byte(elementSize),
	})
}