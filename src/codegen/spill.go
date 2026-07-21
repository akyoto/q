package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/types"
)

// findTempRegister finds a temporary physical register that is not in use.
func (f *Function) findTempRegister(liveSteps []*Step) cpu.Register {
	usedRegisters := bitSet(0)

	for _, live := range liveSteps {
		if live.Register == -1 {
			continue
		}

		usedRegisters.Set(live.Register)
	}

	for _, reg := range f.CPU.General {
		if !usedRegisters.Has(reg) {
			return reg
		}
	}

	panic("no free registers for temporary")
}

// isSpilled returns true if it's a virtual register.
func (f *Function) isSpilled(reg cpu.Register) bool {
	return reg >= f.CPU.MaxRegisters
}

// loadSpill loads a value from a virtual register into a physical register.
func (f *Function) loadSpill(step *Step, destination cpu.Register) {
	f.Assembler.Append(&asm.LoadFixedOffset{
		Index:       f.spillOffset(step.Register),
		Base:        f.CPU.StackPointer,
		Destination: destination,
		Length:      8,
		Scale:       false,
		Signed:      types.IsSigned(step.Value.Type()),
	})
}

// move moves source register to destination register and handles virtual registers.
func (f *Function) move(destinationStep *Step, sourceStep *Step, step *Step) {
	source := sourceStep.Register
	destination := destinationStep.Register

	if source == destination {
		return
	}

	sourceIsSpilled := f.isSpilled(source)
	destinationIsSpilled := f.isSpilled(destination)

	switch {
	case !sourceIsSpilled && !destinationIsSpilled:
		f.Assembler.Append(&asm.Move{
			Destination: destination,
			Source:      source,
		})

	case sourceIsSpilled && !destinationIsSpilled:
		f.loadSpill(sourceStep, destination)

	case !sourceIsSpilled && destinationIsSpilled:
		f.storeSpill(destinationStep, source)

	case sourceIsSpilled && destinationIsSpilled:
		source = f.resolveOperand(sourceStep, step.Live)
		f.storeSpill(destinationStep, source)
	}
}

// resolveOperand returns the register to use for an operand.
// If the operand is spilled, it loads it from the stack first.
func (f *Function) resolveOperand(step *Step, liveSteps []*Step) cpu.Register {
	if !f.isSpilled(step.Register) {
		return step.Register
	}

	tmp := f.findTempRegister(liveSteps)
	f.loadSpill(step, tmp)
	return tmp
}

// spillOffset returns the stack offset in bytes for a virtual register.
func (f *Function) spillOffset(reg cpu.Register) int {
	return int(reg-f.CPU.MaxRegisters) * 8
}

// storeSpill stores a value from a physical register to the spilled stack slot.
func (f *Function) storeSpill(step *Step, source cpu.Register) {
	f.Assembler.Append(&asm.StoreFixedOffset{
		Index:  f.spillOffset(step.Register),
		Base:   f.CPU.StackPointer,
		Source: source,
		Length: 8,
		Scale:  false,
	})
}

// storeSpillNumber stores a number on the stack.
func (f *Function) storeSpillNumber(step *Step, typ types.Type, number int) {
	if typ == types.AnyInt {
		typ = types.Int
	}

	unsigned := types.IsUnsigned(typ)

	if f.build.Arch == config.X86 && (!unsigned && cpu.SizeInt(number) <= 4 || unsigned && cpu.SizeUint(number) <= 4) {
		f.Assembler.Append(&asm.StoreFixedOffsetNumber{
			Index:  f.spillOffset(step.Register),
			Base:   f.CPU.StackPointer,
			Number: int32(number),
			Length: byte(typ.Size()),
			Scale:  false,
		})

		return
	}

	tmp := f.findTempRegister(step.Live)

	f.Assembler.Append(&asm.MoveNumber{
		Destination: tmp,
		Number:      number,
	})

	f.Assembler.Append(&asm.StoreFixedOffset{
		Index:  f.spillOffset(step.Register),
		Base:   f.CPU.StackPointer,
		Source: tmp,
		Length: byte(typ.Size()),
		Scale:  false,
	})
}