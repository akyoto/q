package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// markAlive marks the `live` value in the `instructions` as alive and recursively
// proceeds in the predecessors of `block` if they can reach the definition.
func (f *Function) markAlive(live *step, instructions []ssa.Value, block *ssa.Block) {
	for _, current := range slices.Backward(instructions) {
		currentStep := f.ValueToStep[current]

		if slices.Contains(currentStep.Live, live) {
			return
		}

		currentStep.Live = append(currentStep.Live, live)

		if live.Value == current {
			_, isParam := current.(*ssa.Parameter)

			if !isParam {
				return
			}
		}
	}

	traversed := make(map[*ssa.Block]bool)
	traversed[block] = true

	for _, pre := range block.Predecessors {
		if pre.CanReachPredecessor2(live.Block, traversed) {
			f.markAlive(live, pre.Instructions, pre)
		}
	}
}