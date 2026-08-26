package optimizer

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// Fold evaluates binary operations at compile time and replaces them with constants.
func Fold(ir ssa.IR) map[ssa.Value]struct{} {
	var folded map[ssa.Value]struct{}

	for _, block := range ir.Blocks {
		for _, value := range block.Instructions {
			switch op := value.(type) {
			case *ssa.BinaryOp:
				folded = foldBinaryOp(ir, block, op, folded)
			case *ssa.UnaryOp:
				folded = foldUnaryOp(ir, block, op, folded)
			}
		}
	}

	return folded
}