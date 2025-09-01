package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// markAlive marks the `live` value in the `block` as alive and recursively
// proceeds in the predecessors of `block` if they can reach the definition.
func (f *Function) markAlive(live *Step, block *ssa.Block, use *Step) {
	if use.Block == block {
		phi, isPhi := use.Value.(*ssa.Phi)

		if isPhi {
			index := phi.Arguments.Index(live.Value)
			pre := block.Predecessors[index]
			f.markAlive(live, pre, use)
			return
		}
	}

	region := f.BlockToRegion[block]

	if use.Block == block && (block.Loop == nil || live.Block.Loop != nil) {
		region.End = uint32(use.Index)
	}

	steps := f.Steps[region.Start:region.End]

	for _, current := range slices.Backward(steps) {
		if slices.Contains(current.Live, live) {
			return
		}

		current.Live = append(current.Live, live)

		if live.Value == current.Value {
			_, isParam := current.Value.(*ssa.Parameter)
			_, isPhi := current.Value.(*ssa.Phi)

			if !isParam && !isPhi {
				return
			}
		}
	}

	for _, pre := range block.Predecessors {
		if pre == block {
			continue
		}

		if !pre.CanReachPredecessor(live.Block) {
			continue
		}

		f.markAlive(live, pre, use)
	}
}