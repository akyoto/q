package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeStore(instr *ssa.Store) {
	memory := instr.Memory.(*ssa.Memory)
	address := f.ValueToStep[memory.Address]
	index := f.ValueToStep[memory.Index]
	source := f.ValueToStep[instr.Value]

	if source.Register == -1 {
		f.Assembler.Append(&asm.StoreNumber{
			Base:   address.Register,
			Index:  index.Register,
			Number: source.Value.(*ssa.Int).Int,
			Scale:  memory.Scale,
			Length: byte(memory.Typ.Size()),
		})
	} else {
		f.Assembler.Append(&asm.Store{
			Base:   address.Register,
			Index:  index.Register,
			Source: source.Register,
			Scale:  memory.Scale,
			Length: byte(memory.Typ.Size()),
		})
	}
}