package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// lintBinary checks for some common mistakes in binary expressions.
func (f *Function) lintBinaryOp(binOp *ssa.BinaryOp) error {
	switch binOp.Op {
	case token.Sub, token.Div, token.Mod, token.And, token.Or, token.Xor, token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
		if binOp.Left == binOp.Right {
			return errors.New(&IdenticalExpressions{Operator: binOp.Op.String()}, f.File, binOp.Source.Start())
		}
	}

	if !binOp.Op.IsComparison() {
		return nil
	}

	if types.IsUnsigned(binOp.Left.Type()) {
		rightInt, rightIsInt := binOp.Right.(*ssa.Int)

		if rightIsInt && rightInt.Int < 0 {
			switch binOp.Op {
			case token.Equal, token.LessEqual, token.Less:
				return errors.New(AlwaysFalse, f.File, binOp.Source.Start())
			case token.NotEqual, token.GreaterEqual, token.Greater:
				return errors.New(AlwaysTrue, f.File, binOp.Source.Start())
			}
		}
	}

	if types.IsUnsigned(binOp.Right.Type()) {
		leftInt, leftIsInt := binOp.Left.(*ssa.Int)

		if leftIsInt && leftInt.Int < 0 {
			switch binOp.Op {
			case token.Equal, token.GreaterEqual, token.Greater:
				return errors.New(AlwaysFalse, f.File, binOp.Source.Start())
			case token.NotEqual, token.LessEqual, token.Less:
				return errors.New(AlwaysTrue, f.File, binOp.Source.Start())
			}
		}
	}

	return nil
}