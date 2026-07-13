package linter

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/ssa"
)

// LintBinaryOps checks for some common mistakes in binary expressions.
func LintBinaryOps(ir ssa.IR, file *fs.File) error {
	for _, block := range ir.Blocks {
		for _, value := range block.Instructions {
			binOp, isBinOp := value.(*ssa.BinaryOp)

			if !isBinOp {
				continue
			}

			err := lintBinaryOp(binOp, file)

			if err != nil {
				return err
			}
		}
	}

	return nil
}