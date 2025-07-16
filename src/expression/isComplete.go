package expression

import "git.urbach.dev/cli/q/src/token"

// isComplete returns true if the expression is complete (a binary operation with a single operand is incomplete).
func isComplete(expr *Expression) bool {
	if expr == nil {
		return false
	}

	switch expr.Token.Kind {
	case token.Identifier, token.Number, token.String:
		// These aren't operators but they always count as complete expressions.
		return true
	case token.Call:
		// Even though token.Call is an operator and could be handled by the upcoming branch,
		// the number of operands is variable.
		// Therefore we consider every single call expression as complete.
		return true
	}

	if expr.Token.Kind.IsOperator() && len(expr.Children) == numOperands(expr.Token.Kind) {
		return true
	}

	return false
}