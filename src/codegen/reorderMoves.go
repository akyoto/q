package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/set"
)

// reorderMoves reorders the move instructions.
func reorderMoves(moves []asm.Instruction) {
	usedRegisters := bitSet(0)
	futureRegisters := bitSet(0)

	for i, instr := range moves {
		var (
			source         cpu.Register
			destination    cpu.Register
			hasSource      bool
			hasDestination bool
		)

		switch instr := instr.(type) {
		case *asm.Move:
			source = instr.Source
			destination = instr.Destination
			hasSource = true
			hasDestination = true
		case *asm.StoreFixedOffset:
			source = instr.Source
			hasSource = true
		case *asm.LoadFixedOffset:
			destination = instr.Destination
			hasDestination = true
		default:
			continue
		}

		if hasSource && futureRegisters.Has(source) {
			set.BringToFront(moves[:i+1], i)

			if hasDestination && usedRegisters.Has(destination) {
				panic("cycle detected while reordering moves")
			}
		}

		if hasSource {
			usedRegisters.Set(source)
		}

		if hasDestination {
			futureRegisters.Set(destination)
		}
	}
}