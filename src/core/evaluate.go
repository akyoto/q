package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluate converts an expression to an SSA value.
func (f *Function) evaluate(expr *expression.Expression) (ssa.Value, error) {
	if expr.IsLeaf() {
		return f.evaluateLeaf(expr)
	}

	switch expr.Token.Kind {
	case token.Call:
		return f.evaluateCall(expr)

	case token.Dot:
		return f.evaluateDot(expr)

	case token.Array:
		return f.evaluateArray(expr)

	case token.Struct:
		return f.evaluateStruct(expr)

	case token.Cast:
		return f.evaluateCast(expr)
	}

	if len(expr.Children) == 1 {
		return f.evaluateUnary(expr)
	}

	return f.evaluateBinary(expr)
}