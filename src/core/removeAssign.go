package core

import "git.urbach.dev/cli/q/src/token"

// removeAssign removes the assignment part of an operator token and returns the raw operation.
func removeAssign(kind token.Kind) token.Kind {
	switch kind {
	case token.AddAssign:
		return token.Add
	case token.SubAssign:
		return token.Sub
	case token.MulAssign:
		return token.Mul
	case token.DivAssign:
		return token.Div
	case token.ModAssign:
		return token.Mod
	case token.ShlAssign:
		return token.Shl
	case token.ShrAssign:
		return token.Shr
	case token.AndAssign:
		return token.And
	case token.OrAssign:
		return token.Or
	case token.XorAssign:
		return token.Xor
	default:
		panic("not implemented")
	}
}