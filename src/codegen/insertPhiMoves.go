package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
)

// insertPhiMoves moves all live values that are part of a Phi instruction
// from their current register to the Phi target register.
// It must be called right before a Jump instruction.
func (f *Function) insertPhiMoves(step *Step) {
	for _, live := range step.Live {
		for phi := range live.Phis.All() {
			if live.Register == phi.Register {
				continue
			}

			if !slices.Contains(phi.Block.Predecessors, step.Block) {
				continue
			}

			f.Assembler.Append(&asm.Move{
				Destination: phi.Register,
				Source:      live.Register,
			})
		}
	}
}