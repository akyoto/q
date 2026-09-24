package optimizer

import "git.urbach.dev/cli/q/src/ssa"

// foldTo replaces a binary operation with another value.
func foldTo(ir ssa.IR, binaryOp *ssa.BinaryOp, value ssa.Value, folded map[ssa.Value]struct{}) map[ssa.Value]struct{} {
	if folded == nil {
		folded = make(map[ssa.Value]struct{})
	}

	folded[binaryOp] = struct{}{}
	ir.ReplaceAll(binaryOp, value)
	return folded
}