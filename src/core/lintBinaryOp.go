package core

import (
	"strings"

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

	leftType := binOp.Left.Type()
	rightType := binOp.Right.Type()
	leftIsUnsigned := types.IsUnsigned(leftType)
	rightIsUnsigned := types.IsUnsigned(rightType)

	switch binOp.Op {
	case token.Div, token.Mod, token.Shr, token.Greater, token.GreaterEqual, token.Less, token.LessEqual:
		if leftIsUnsigned && !rightIsUnsigned && rightType != types.AnyInt {
			return errors.New(&MixedSignedUnsigned{Signed: rightType.Name(), Unsigned: leftType.Name()}, f.File, binOp.Source)
		}

		if !leftIsUnsigned && rightIsUnsigned && leftType != types.AnyInt {
			return errors.New(&MixedSignedUnsigned{Signed: leftType.Name(), Unsigned: rightType.Name()}, f.File, binOp.Source)
		}
	}

	if !binOp.Op.IsComparison() {
		return nil
	}

	if leftIsUnsigned {
		rightInt, rightIsInt := binOp.Right.(*ssa.Int)

		if rightIsInt {
			switch {
			case rightInt.Int == 0:
				switch binOp.Op {
				case token.Less:
					return errors.New(AlwaysFalse, f.File, binOp.Source)
				case token.GreaterEqual:
					return errors.New(AlwaysTrue, f.File, binOp.Source)
				}

			case rightInt.Int < 0 && strings.HasPrefix(rightInt.StringFrom(f.File.Bytes), "-"):
				switch binOp.Op {
				case token.Equal, token.LessEqual, token.Less:
					return errors.New(AlwaysFalse, f.File, binOp.Source)
				case token.NotEqual, token.GreaterEqual, token.Greater:
					return errors.New(AlwaysTrue, f.File, binOp.Source)
				}
			}
		}
	}

	if rightIsUnsigned {
		leftInt, leftIsInt := binOp.Left.(*ssa.Int)

		if leftIsInt {
			switch {
			case leftInt.Int == 0:
				switch binOp.Op {
				case token.Greater:
					return errors.New(AlwaysFalse, f.File, binOp.Source)
				case token.LessEqual:
					return errors.New(AlwaysTrue, f.File, binOp.Source)
				}

			case leftInt.Int < 0 && strings.HasPrefix(leftInt.StringFrom(f.File.Bytes), "-"):
				switch binOp.Op {
				case token.Equal, token.GreaterEqual, token.Greater:
					return errors.New(AlwaysFalse, f.File, binOp.Source)
				case token.NotEqual, token.LessEqual, token.Less:
					return errors.New(AlwaysTrue, f.File, binOp.Source)
				}
			}
		}
	}

	return nil
}