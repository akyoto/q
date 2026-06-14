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

			isAssociative := binaryOp.IsAssociative()
			foldLeft := binaryOp.Left
			leftBinOp, leftIsBinOp := foldLeft.(*ssa.BinaryOp)

			if isAssociative && leftIsBinOp && leftBinOp.Op == binaryOp.Op {
				foldLeft = leftBinOp.Right
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
				binaryOp.Left = leftBinOp.Left
				binaryOp.Right = constant
				block.InsertAt(i, constant)

			case !leftIsBinOp && rightIsBinOp:
				folded[rightBinOp] = struct{}{}
				binaryOp.Left = constant
				binaryOp.Right = rightBinOp.Right
				block.InsertAt(i, constant)
			}
		}
	}

	return folded
}