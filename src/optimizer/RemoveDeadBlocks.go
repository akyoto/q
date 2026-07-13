package optimizer

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// RemoveDeadBlocks removes unreachable blocks.
func RemoveDeadBlocks(ir *ssa.IR) {
	ir.Blocks = slices.DeleteFunc(ir.Blocks, func(block *ssa.Block) bool {
		return len(block.Predecessors) == 0 && block != ir.Blocks[0]
	})
}