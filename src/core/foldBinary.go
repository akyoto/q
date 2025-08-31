package core

import (
	"git.urbach.dev/cli/q/src/token"
)

// foldBinary calculates the result of the binary operation.
func foldBinary(op token.Kind, a int, b int) int {
	switch op {
	case token.Add:
		return a + b
	case token.And:
		return a & b
	case token.Div:
		return a / b
	case token.Mul:
		return a * b
	case token.Mod:
		return a % b
	case token.Or:
		return a | b
	case token.Shl:
		return a << b
	case token.Shr:
		return a >> b
	case token.Sub:
		return a - b
	case token.Xor:
		return a ^ b
	default:
		panic("unknown fold operation: " + op.String())
	}
}