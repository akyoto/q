package optimizer

import (
	"math/bits"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// foldPowerOfTwo rewrites multiplication, division and modulo when used with a power of 2.
func foldPowerOfTwo(block *ssa.Block, op *ssa.BinaryOp, folded map[ssa.Value]struct{}) map[ssa.Value]struct{} {
	if folded != nil {
		_, done := folded[op]

		if done {
			return folded
		}
	}

	var (
		newOp     token.Kind
		newValue  int
		discarded ssa.Value
	)

	switch op.Op {
	case token.Mul:
		n, isPowerOfTwo := exponent(op.Right)

		if isPowerOfTwo {
			newOp = token.Shl
			newValue = n
			discarded = op.Right
		} else {
			n, isPowerOfTwo = exponent(op.Left)

			if !isPowerOfTwo {
				return folded
			}

			newOp = token.Shl
			newValue = n
			discarded = op.Left
			op.Left = op.Right
		}
	case token.Div, token.Mod:
		n, isPowerOfTwo := exponent(op.Right)

		if !isPowerOfTwo || !types.IsUnsigned(op.Left.Type()) {
			return folded
		}

		discarded = op.Right

		if op.Op == token.Div {
			newOp = token.Shr
			newValue = n
		} else {
			newOp = token.And
			newValue = (1 << n) - 1
		}
	default:
		return folded
	}

	i := block.Index(op)

	if i == -1 {
		return folded
	}

	value := &ssa.Int{Int: newValue}
	block.InsertAt(i, value)

	if folded == nil {
		folded = make(map[ssa.Value]struct{})
	}

	folded[discarded] = struct{}{}
	op.Op = newOp
	op.Right = value
	op.Swapped = false
	return folded
}

// exponent returns the base 2 exponent of a positive power of 2.
func exponent(value ssa.Value) (int, bool) {
	number, isInt := value.(*ssa.Int)

	if !isInt || number.Int < 2 || number.Int&(number.Int-1) != 0 {
		return 0, false
	}

	return bits.TrailingZeros64(uint64(number.Int)), true
}