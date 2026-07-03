package optimizer

import "git.urbach.dev/cli/q/src/ssa"

// foldUnaryOp folds the values for unary operations.
func foldUnaryOp(ir ssa.IR, block *ssa.Block, i int, unaryOp *ssa.UnaryOp, folded map[ssa.Value]struct{}) map[ssa.Value]struct{} {
	left, leftIsInt := unaryOp.Operand.(*ssa.Int)

	if !leftIsInt {
		return folded
	}

	if folded == nil {
		folded = make(map[ssa.Value]struct{})
	}

	folded[left] = struct{}{}

	constant := &ssa.Int{
		Int:    foldUnary(unaryOp.Op, left.Int),
		Source: unaryOp.Source,
	}

	block.Instructions[i] = constant
	ir.ReplaceAll(unaryOp, constant)
	return folded
}