package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// lintBinaryOps checks for some common mistakes in binary expressions.
func (f *Function) lintBinaryOps() error {
	for _, block := range f.Blocks {
		for _, value := range block.Instructions {
			binOp, isBinOp := value.(*ssa.BinaryOp)

			if !isBinOp {
				continue
			}

			err := f.lintBinaryOp(binOp)

			if err != nil {
				return err
			}
		}
	}

	return nil
}