package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// NewList generates a list of expressions from comma separated parameters.
func NewList(tokens token.List) []*Expression {
	var list []*Expression

	for position, param := range tokens.Split {
		expression := Parse(param)

		if expression.Token.Kind == token.Invalid {
			expression.Token.Position = position
		}

		list = append(list, expression)
	}

	return list
}