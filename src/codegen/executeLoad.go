package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

func (f *Function) executeLoad(step *Step, instr *ssa.Load) {
	if step.Register == -1 {
		return
	}

	memory := instr.Memory
	address := f.ValueToStep[memory.Address]
	index := f.ValueToStep[memory.Index]
	elementType := step.Value.Type()
	elementSize := elementType.Size()
	baseRegister := f.resolveOperand(address, step.Live)
	indexRegister := f.resolveOperand(index, step.Live)
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	if index.Register == -1 {
		fixedOffset := index.Value.(*ssa.Int).Int
		f.Assembler.Append(&asm.LoadFixedOffset{
			Base:        baseRegister,
			Index:       fixedOffset,
			Destination: destination,
			Scale:       memory.Scale,
			Length:      byte(elementSize),
			Signed:      types.IsSigned(elementType),
		})
	} else {
		f.Assembler.Append(&asm.Load{
			Base:        baseRegister,
			Index:       indexRegister,
			Destination: destination,
			Scale:       memory.Scale,
			Length:      byte(elementSize),
			Signed:      types.IsSigned(elementType),
		})
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}