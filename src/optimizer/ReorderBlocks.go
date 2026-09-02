package optimizer

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// ReorderBlocks moves rarely taken branch bodies to the end of the function.
func ReorderBlocks(ir *ssa.IR) {
	blocks := ir.Blocks

	for i := 0; i < len(blocks)-1; i++ {
		if blocks[i].Loop != nil {
			continue
		}

		branch, isBranch := blocks[i].Last().(*ssa.Branch)

		if !isBranch {
			continue
		}

		rare := blocks[i+1]

		if !isRareBody(blocks[i], rare) {
			continue
		}

		normal := branch.Then

		if rare == normal {
			normal = branch.Else
		}

		end := -1

		for j := i + 1; j < len(blocks); j++ {
			if blocks[j] == normal {
				end = j
				break
			}
		}

		if end == -1 {
			continue
		}

		rareRun := slices.Clone(blocks[i+1 : end])
		blocks = append(blocks[:i+1], blocks[end:]...)
		blocks = append(blocks, rareRun...)
	}

	ir.Blocks = blocks
}

// isRareBody checks if the body can only be reached from the given block.
func isRareBody(block *ssa.Block, body *ssa.Block) bool {
	if len(body.Predecessors) != 1 || body.Predecessors[0] != block {
		return false
	}

	switch last := body.Last().(type) {
	case *ssa.Call:
		return last.Func.Typ.NoReturn
	case *ssa.Jump, *ssa.Return:
		return true
	default:
		return false
	}
}