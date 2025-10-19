package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// lintBinaryOp checks for some common mistakes in binary expressions.
func (f *Function) lintBinaryOp(binOp *ssa.BinaryOp) error {
	if binOp.Source.End() == 0 {
		return nil
	}

	if binOp.Left == binOp.Right {
		switch binOp.Op {
		case token.Sub, token.Div, token.Mod, token.And, token.Or, token.Xor, token.Equal, token.NotEqual, token.Less, token.LessEqual, token.Greater, token.GreaterEqual:
			return errors.New(&IdenticalExpressions{Operator: binOp.Op.String()}, f.File, binOp.Source)
		}
	}

	left, leftIsInt := binOp.Left.(*ssa.Int)
	right, rightIsInt := binOp.Right.(*ssa.Int)

	switch binOp.Op {
	case token.Add, token.Sub:
		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: binOp.Right.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

	case token.And:
		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

	case token.Mul:
		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

		if leftIsInt && left.Int == 1 {
			return errors.New(&Simplify{To: binOp.Right.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 1 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

	case token.Div:
		if rightIsInt && right.Int == 0 {
			return errors.New(DivisionByZero, f.File, binOp.Source)
		}

		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 1 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

	case token.Mod:
		if rightIsInt && right.Int == 0 {
			return errors.New(DivisionByZero, f.File, binOp.Source)
		}

		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 1 {
			return errors.New(&Simplify{To: "0"}, f.File, binOp.Source)
		}

	case token.Or:
		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: binOp.Right.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

	case token.Shl, token.Shr:
		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

	case token.Xor:
		if leftIsInt && left.Int == 0 {
			return errors.New(&Simplify{To: binOp.Right.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
		}

		if rightIsInt && right.Int == 0 {
			return errors.New(&Simplify{To: binOp.Left.(errors.Source).StringFrom(f.File.Bytes)}, f.File, binOp.Source)
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
				return errors.New(AlwaysFalse, f.File, binOp.Source)
			case token.NotEqual, token.GreaterEqual, token.Greater:
				return errors.New(AlwaysTrue, f.File, binOp.Source)
			}
		}
	}

	if types.IsUnsigned(binOp.Right.Type()) {
		leftInt, leftIsInt := binOp.Left.(*ssa.Int)

		if leftIsInt && leftInt.Int < 0 {
			switch binOp.Op {
			case token.Equal, token.GreaterEqual, token.Greater:
				return errors.New(AlwaysFalse, f.File, binOp.Source)
			case token.NotEqual, token.LessEqual, token.Less:
				return errors.New(AlwaysTrue, f.File, binOp.Source)
			}
		}
	}

	return nil
}