package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

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
		Signed:      !types.IsUnsigned(step.Value.Type()),
	})
}

// spillOffset returns the stack offset in bytes for a virtual register.
func (f *Function) spillOffset(reg cpu.Register) int {
	return int(reg-f.CPU.MaxRegisters) * 8
}

// storeSpillNumber stores a number on the stack.
func (f *Function) storeSpillNumber(step *Step, number *ssa.Int) {
	typ := number.Type()

	if typ == types.AnyInt {
		typ = types.Int
	}

	unsigned := types.IsUnsigned(typ)

	if f.build.Arch == config.X86 && (!unsigned && cpu.SizeInt(number.Int) <= 4 || unsigned && cpu.SizeUint(number.Int) <= 4) {
		f.Assembler.Append(&asm.StoreFixedOffsetNumber{
			Index:  f.spillOffset(step.Register),
			Base:   f.CPU.StackPointer,
			Number: int32(number.Int),
			Length: byte(typ.Size()),
			Scale:  false,
		})

		return
	}

	usedRegisters := bitSet(0)
	tmp := cpu.Register(0)

	for _, live := range step.Live {
		if live.Register == -1 || f.isSpilled(live.Register) {
			continue
		}

		usedRegisters.Set(live.Register)
	}

	for _, reg := range f.CPU.General {
		if !usedRegisters.Has(reg) {
			tmp = reg
			break
		}
	}

	f.Assembler.Append(&asm.MoveNumber{
		Destination: tmp,
		Number:      number.Int,
	})

	f.Assembler.Append(&asm.StoreFixedOffset{
		Index:  f.spillOffset(step.Register),
		Base:   f.CPU.StackPointer,
		Source: tmp,
		Length: byte(typ.Size()),
		Scale:  false,
	})
}