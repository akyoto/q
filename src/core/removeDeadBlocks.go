package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// removeDeadBlocks removes unreachable blocks.
func (f *Function) removeDeadBlocks() {
	f.IR.Blocks = slices.DeleteFunc(f.IR.Blocks, func(block *ssa.Block) bool {
		return len(block.Predecessors) == 0 && block != f.IR.Blocks[0]
	})
}