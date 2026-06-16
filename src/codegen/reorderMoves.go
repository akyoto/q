package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/set"
)

// reorderMoves reorders the move instructions.
func reorderMoves(moves []asm.Instruction) {
	usedRegisters := bitSet(0)
	futureRegisters := bitSet(0)

	for i, instr := range moves {
		move := instr.(*asm.Move)

		if futureRegisters.Has(move.Source) {
			set.BringToFront(moves[:i+1], i)

			if usedRegisters.Has(move.Destination) {
				panic("cycle detected while reordering moves")
			}
		}

		usedRegisters.Set(move.Source)
		futureRegisters.Set(move.Destination)
	}
}