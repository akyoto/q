package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluateBuiltin converts a call to a builtin function to an SSA value.
func (f *Function) evaluateBuiltin(expr *expression.Expression) (ssa.Value, error) {
	switch expr.Children[0].Token.Kind {
	case token.New:
		return f.evaluateNew(expr)

	case token.Delete:
		return f.evaluateDelete(expr)

	case token.Syscall:
		return f.evaluateSyscall(expr)

	default:
		panic("not implemented")
	}
}