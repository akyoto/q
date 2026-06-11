package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// isTrailing is used for trailing commas and returns true if an expression
// has an invalid token and is the last expression in the list.
func isTrailing(definition *expression.Expression, siblings []*expression.Expression) bool {
	if definition.Token.Kind != token.Invalid {
		return false
	}

	return definition == siblings[len(siblings)-1]
}