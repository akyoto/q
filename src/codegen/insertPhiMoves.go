package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// phiMove holds temporary data for scheduling moves to Phi target registers.
type phiMove struct {
	Destination *Step
	Source      *Step
	IsSwap      bool
	IsHandled   bool
}

// insertPhiMoves moves all live values that are part of a Phi instruction
// from their current register to the Phi target register.
// It must be called right before a Jump instruction.
func (f *Function) insertPhiMoves(step *Step) {
	var phiMoves []phiMove

	for _, live := range step.Live {
		for phi := range live.Phis.All() {
			predecessors := phi.Block.Predecessors

			if !slices.Contains(predecessors, step.Block) {
				continue
			}

			matches := false

			for index, value := range phi.Value.(*ssa.Phi).Arguments {
				if value == live.Value && predecessors[index] == step.Block {
					matches = true
					break
				}
			}

			if !matches {
				continue
			}

			phiMoves = append(phiMoves, phiMove{
				Destination: phi,
				Source:      live,
			})
		}
	}

	for i, current := range phiMoves {
		source := current.Source.Register
		destination := current.Destination.Register

		if current.IsSwap || f.isSpilled(source) || f.isSpilled(destination) {
			continue
		}

		for j, other := range phiMoves {
			if j == i {
				continue
			}

			if other.Source.Register == destination && other.Destination.Register == source {
				phiMoves[i].IsSwap = true
				phiMoves[j].IsSwap = true
				break
			}
		}
	}

	start := len(f.Assembler.Instructions)

	for i, move := range phiMoves {
		if move.IsSwap {
			continue
		}

		f.move(move.Destination, move.Source, step)
		phiMoves[i].IsHandled = true
	}

	end := len(f.Assembler.Instructions)
	reorderMoves(f.Assembler.Instructions[start:end])

	for i, move := range phiMoves {
		if !move.IsSwap || move.IsHandled {
			continue
		}

		source := move.Source.Register
		destination := move.Destination.Register
		tmp := f.findTempRegister(step.Live)
		f.Assembler.Append(&asm.Move{Destination: tmp, Source: source})
		f.Assembler.Append(&asm.Move{Destination: source, Source: destination})
		f.Assembler.Append(&asm.Move{Destination: destination, Source: tmp})

		for j, other := range phiMoves {
			if other.Source.Register == destination && other.Destination.Register == source {
				phiMoves[j].IsHandled = true
			}
		}

		phiMoves[i].IsHandled = true
	}
}