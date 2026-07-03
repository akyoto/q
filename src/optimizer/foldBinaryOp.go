package optimizer

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// foldBinaryOp folds the values for binary operations.
func foldBinaryOp(ir ssa.IR, block *ssa.Block, i int, binaryOp *ssa.BinaryOp, folded map[ssa.Value]struct{}) map[ssa.Value]struct{} {
	if binaryOp.Op.IsComparison() {
		return folded
	}

	isAssociative := binaryOp.Op.IsAssociative()
	foldLeft := binaryOp.Left
	leftBinOp, leftIsBinOp := foldLeft.(*ssa.BinaryOp)

	if isAssociative && leftIsBinOp && leftBinOp.Op == binaryOp.Op {
		innerRight, innerRightIsInt := leftBinOp.Right.(*ssa.Int)

		if innerRightIsInt {
			foldLeft = innerRight
		} else {
			if !binaryOp.Op.IsCommutative() {
				return folded
			}

			innerLeft, innerLeftIsInt := leftBinOp.Left.(*ssa.Int)

			if !innerLeftIsInt {
				return folded
			}

			foldLeft = innerLeft
		}
	}

	left, leftIsInt := foldLeft.(*ssa.Int)

	if !leftIsInt {
		return folded
	}

	foldRight := binaryOp.Right
	rightBinOp, rightIsBinOp := foldRight.(*ssa.BinaryOp)

	if isAssociative && rightIsBinOp && rightBinOp.Op == binaryOp.Op {
		foldRight = rightBinOp.Left
	}

	if isAssociative && rightIsBinOp && rightBinOp.Op == binaryOp.Op {
		innerLeft, innerLeftIsInt := rightBinOp.Left.(*ssa.Int)

		if innerLeftIsInt {
			foldRight = innerLeft
		} else {
			if !binaryOp.Op.IsCommutative() {
				return folded
			}

			innerRight, innerRightIsInt := rightBinOp.Right.(*ssa.Int)

			if !innerRightIsInt {
				return folded
			}

			foldRight = innerRight
		}
	}

	right, rightIsInt := foldRight.(*ssa.Int)

	if !rightIsInt {
		return folded
	}

	if (binaryOp.Op == token.Div || binaryOp.Op == token.Mod) && right.Int == 0 {
		return folded
	}

	if leftIsBinOp && rightIsBinOp {
		return folded
	}

	if folded == nil {
		folded = make(map[ssa.Value]struct{})
	}

	folded[foldLeft] = struct{}{}
	folded[foldRight] = struct{}{}

	constant := &ssa.Int{
		Int:    foldBinary(binaryOp.Op, left.Int, right.Int),
		Source: binaryOp.Source,
	}

	switch {
	case !leftIsBinOp && !rightIsBinOp:
		block.Instructions[i] = constant
		ir.ReplaceAll(binaryOp, constant)

	case leftIsBinOp && !rightIsBinOp:
		folded[leftBinOp] = struct{}{}

		if foldLeft == leftBinOp.Right {
			binaryOp.Left = leftBinOp.Left
		} else {
			binaryOp.Left = leftBinOp.Right
		}

		binaryOp.Right = constant
		block.InsertAt(i, constant)

	case !leftIsBinOp && rightIsBinOp:
		folded[rightBinOp] = struct{}{}
		binaryOp.Left = constant

		if foldRight == rightBinOp.Left {
			binaryOp.Right = rightBinOp.Right
		} else {
			binaryOp.Right = rightBinOp.Left
		}

		block.InsertAt(i, constant)
	}

	return folded
}