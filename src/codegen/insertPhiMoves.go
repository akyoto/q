package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// insertPhiMoves moves all live values that are part of a Phi instruction
// from their current register to the Phi target register.
// It must be called right before a Jump instruction.
func (f *Function) insertPhiMoves(step *Step) {
	var (
		moves            []*asm.Move
		sourceSteps      = map[cpu.Register]*Step{}
		destinationSteps = map[cpu.Register]*Step{}
	)

	for _, live := range step.Live {
		for phi := range live.Phis.All() {
			if live.Register == phi.Register {
				continue
			}

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

			destinationSteps[phi.Register] = phi
			sourceSteps[live.Register] = live

			moves = append(moves, &asm.Move{
				Destination: phi.Register,
				Source:      live.Register,
			})
		}
	}

	free := f.freeTempRegisters(step.Live)
	scheduled, ok := ScheduleMoves(moves, free)

	if !ok {
		panic("no free register for move scheduling")
	}

	for _, move := range scheduled {
		source := move.Source
		destination := move.Destination

		switch {
		case f.isSpilled(source) && f.isSpilled(destination):
			tmp := f.findTempRegister(step.Live)
			f.loadSpill(sourceSteps[source], tmp)
			f.storeSpill(destinationSteps[destination], tmp)
		case f.isSpilled(source):
			f.loadSpill(sourceSteps[source], destination)
		case f.isSpilled(destination):
			f.storeSpill(destinationSteps[destination], source)
		default:
			f.Assembler.Append(move)
		}
	}
}