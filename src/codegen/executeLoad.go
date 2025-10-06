package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeLoad(step *Step, instr *ssa.Load) {
	if step.Register == -1 {
		return
	}

	memory := instr.Memory
	address := f.ValueToStep[memory.Address]
	index := f.ValueToStep[memory.Index]
	elementSize := step.Value.Type().Size()

	if index.Register == -1 {
		f.Assembler.Append(&asm.LoadFixedOffset{
			Base:        address.Register,
			Index:       index.Value.(*ssa.Int).Int,
			Destination: step.Register,
			Scale:       memory.Scale,
			Length:      byte(elementSize),
		})
	} else {
		f.Assembler.Append(&asm.Load{
			Base:        address.Register,
			Index:       index.Register,
			Destination: step.Register,
			Scale:       memory.Scale,
			Length:      byte(elementSize),
		})
	}
}