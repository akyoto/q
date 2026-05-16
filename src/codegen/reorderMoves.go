package codegen

import "git.urbach.dev/cli/q/src/asm"

// reorderMoves reorders the move instructions.
func reorderMoves(moves []asm.Instruction) {
	usedRegisters := 0
	futureRegisters := 0

	for i, instr := range moves {
		move := instr.(*asm.Move)

		if futureRegisters&(1<<move.Source) != 0 {
			bringToFront(moves[:i+1], i)

			if usedRegisters&(1<<move.Destination) != 0 {
				panic("cycle detected while reordering moves")
			}
		}

		usedRegisters |= (1 << move.Source)
		futureRegisters |= (1 << move.Destination)
	}
}