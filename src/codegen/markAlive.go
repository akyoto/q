package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// markAlive marks the `live` value in the `instructions` as alive and recursively
// proceeds in the predecessors of `block` if they can reach the definition.
func (f *Function) markAlive(live *Step, instructions []ssa.Value, block *ssa.Block, use *Step) {
	for _, current := range slices.Backward(instructions) {
		currentStep := f.ValueToStep[current]

		if slices.Contains(currentStep.Live, live) {
			return
		}

		currentStep.Live = append(currentStep.Live, live)

		if live.Value == current {
			_, isParam := current.(*ssa.Parameter)
			_, isPhi := current.(*ssa.Phi)

			if !isParam && !isPhi {
				return
			}
		}
	}

	if use.Block == block {
		switch instr := use.Value.(type) {
		case *ssa.Phi:
			index := instr.Arguments.Index(live.Value)
			pre := block.Predecessors[index]
			f.markAlive(live, pre.Instructions, pre, use)
			return
		}
	}

	for _, pre := range block.Predecessors {
		if pre.CanReachPredecessor(live.Block) {
			f.markAlive(live, pre.Instructions, pre, use)
		}
	}
}