package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// lintBinary checks for some common mistakes in binary expressions.
func (f *Function) lintBinaryOp(binOp *ssa.BinaryOp) error {
	if !binOp.Op.IsComparison() {
		return nil
	}

	switch binOp.Left.Type() {
	case types.UInt64, types.UInt32, types.UInt16, types.UInt8:
		rightInt, rightIsInt := binOp.Right.(*ssa.Int)

		if rightIsInt && rightInt.Int < 0 {
			var err error

			switch binOp.Op {
			case token.Equal, token.LessEqual, token.Less:
				err = AlwaysFalse
			case token.NotEqual, token.GreaterEqual, token.Greater:
				err = AlwaysTrue
			}

			return errors.New(err, f.File, binOp.Source.Start())
		}
	}

	switch binOp.Right.Type() {
	case types.UInt64, types.UInt32, types.UInt16, types.UInt8:
		leftInt, leftIsInt := binOp.Left.(*ssa.Int)

		if leftIsInt && leftInt.Int < 0 {
			var err error

			switch binOp.Op {
			case token.Equal, token.GreaterEqual, token.Greater:
				err = AlwaysFalse
			case token.NotEqual, token.LessEqual, token.Less:
				err = AlwaysTrue
			}

			return errors.New(err, f.File, binOp.Source.Start())
		}
	}

	return nil
}