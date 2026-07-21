package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeStore(instr *ssa.Store) {
	memory := instr.Memory
	address := f.ValueToStep[memory.Address]
	index := f.ValueToStep[memory.Index]
	source := f.ValueToStep[instr.Value]
	baseRegister := f.resolveOperand(address, source.Live)
	indexRegister := f.resolveOperand(index, source.Live)
	sourceRegister := f.resolveOperand(source, source.Live)

	if sourceRegister == -1 {
		if indexRegister == -1 {
			f.Assembler.Append(&asm.StoreFixedOffsetNumber{
				Base:   baseRegister,
				Index:  index.Value.(*ssa.Int).Int,
				Number: int32(source.Value.(*ssa.Int).Int),
				Scale:  memory.Scale,
				Length: byte(memory.Typ.Size()),
			})
		} else {
			f.Assembler.Append(&asm.StoreNumber{
				Base:   baseRegister,
				Index:  indexRegister,
				Number: int32(source.Value.(*ssa.Int).Int),
				Scale:  memory.Scale,
				Length: byte(memory.Typ.Size()),
			})
		}
	} else {
		if indexRegister == -1 {
			f.Assembler.Append(&asm.StoreFixedOffset{
				Base:   baseRegister,
				Index:  index.Value.(*ssa.Int).Int,
				Source: sourceRegister,
				Scale:  memory.Scale,
				Length: byte(memory.Typ.Size()),
			})
		} else {
			f.Assembler.Append(&asm.Store{
				Base:   baseRegister,
				Index:  indexRegister,
				Source: sourceRegister,
				Scale:  memory.Scale,
				Length: byte(memory.Typ.Size()),
			})
		}
	}
}