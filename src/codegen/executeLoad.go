package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeLoad(step *Step, instr *ssa.Load) {
	if step.Register == -1 {
		return
	}

	memory := instr.Memory.(*ssa.Memory)
	address := f.ValueToStep[memory.Address]
	index := f.ValueToStep[memory.Index]
	elementSize := step.Value.Type().Size()

	f.Assembler.Append(&asm.Load{
		Base:        address.Register,
		Index:       index.Register,
		Destination: step.Register,
		Scale:       memory.Scale,
		Length:      byte(elementSize),
	})
}