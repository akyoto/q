package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// multiplySize multiplies the `count` with the `size` if needed.
func (f *Function) multiplySize(count ssa.Value, size int) ssa.Value {
	if size == 1 {
		return count
	}

	sizeValue := f.Append(&ssa.Int{Int: size})

	return f.Append(&ssa.BinaryOp{
		Op:    token.Mul,
		Left:  count,
		Right: sizeValue,
	})
}