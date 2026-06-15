package optimizer

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Fold evaluates binary operations at compile time and replaces them with constants.
func Fold(ir ssa.IR) map[ssa.Value]struct{} {
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

			isAssociative := binaryOp.Op.IsAssociative()
			foldLeft := binaryOp.Left
			leftBinOp, leftIsBinOp := foldLeft.(*ssa.BinaryOp)

			if isAssociative && leftIsBinOp && leftBinOp.Op == binaryOp.Op {
				innerRight, innerRightIsInt := leftBinOp.Right.(*ssa.Int)

				if innerRightIsInt {
					foldLeft = innerRight
				} else {
					if !binaryOp.Op.IsCommutative() {
						continue
					}

					innerLeft, innerLeftIsInt := leftBinOp.Left.(*ssa.Int)

					if !innerLeftIsInt {
						continue
					}

					foldLeft = innerLeft
				}
			}

			left, leftIsInt := foldLeft.(*ssa.Int)

			if !leftIsInt {
				continue
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
						continue
					}

					innerRight, innerRightIsInt := rightBinOp.Right.(*ssa.Int)

					if !innerRightIsInt {
						continue
					}

					foldRight = innerRight
				}
			}

			right, rightIsInt := foldRight.(*ssa.Int)

			if !rightIsInt {
				continue
			}

			if (binaryOp.Op == token.Div || binaryOp.Op == token.Mod) && right.Int == 0 {
				continue
			}

			if leftIsBinOp && rightIsBinOp {
				continue
			}

			if folded == nil {
				folded = make(map[ssa.Value]struct{})
			}

			folded[foldLeft] = struct{}{}
			folded[foldRight] = struct{}{}

			constant := &ssa.Int{
				Int:    FoldBinary(binaryOp.Op, left.Int, right.Int),
				Source: binaryOp.Source,
			}

			switch {
			case !leftIsBinOp && !rightIsBinOp:
				block.Instructions[i] = constant
				ir.ReplaceAll(value, constant)

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
		}
	}

	return folded
}