package optimizer

import "git.urbach.dev/cli/q/src/token"

// foldUnary calculates the result of the unary operation.
func foldUnary(op token.Kind, a int) int {
	switch op {
	case token.Negate:
		return -a
	default:
		panic("unknown fold operation: " + op.String())
	}
}