package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// isIdentifierRead returns true if the expression reads an identifier.
func isIdentifierRead(expr *expression.Expression) bool {
	return expr.Token.Kind == token.Identifier || expr.Token.Kind == token.Dot
}