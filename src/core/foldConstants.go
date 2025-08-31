package core

import "git.urbach.dev/cli/q/src/ssa"

// foldConstants evaluates binary operations at compile time and replaces them with constants.
func foldConstants(ir ssa.IR) map[ssa.Value]struct{} {
	var folded map[ssa.Value]struct{}

	for _, block := range ir.Blocks {
		for i, value := range block.Instructions {
			binaryOp, isBinaryOp := value.(*ssa.BinaryOp)

			if !isBinaryOp {
				continue
			}

			if binaryOp.Op.IsComparison() {
				continue
			}

			left, leftIsInt := binaryOp.Left.(*ssa.Int)

			if !leftIsInt {
				continue
			}

			right, rightIsInt := binaryOp.Right.(*ssa.Int)

			if !rightIsInt {
				continue
			}

			constant := &ssa.Int{
				Int:    foldBinary(binaryOp.Op, left.Int, right.Int),
				Source: binaryOp.Source,
			}

			if folded == nil {
				folded = make(map[ssa.Value]struct{})
			}

			folded[binaryOp.Left] = struct{}{}
			folded[binaryOp.Right] = struct{}{}
			block.Instructions[i] = constant
			ir.ReplaceAll(value, constant)
		}
	}

	return folded
}