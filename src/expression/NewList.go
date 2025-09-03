package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// NewList generates a list of expressions from comma separated parameters.
func NewList(tokens token.List) []*Expression {
	if len(tokens) == 0 {
		return nil
	}

	list := make([]*Expression, 0, 4)

	for position, param := range tokens.Split {
		expression := Parse(param)

		if expression.Token.Kind == token.Invalid {
			expression.Token.Position = position
		}

		list = append(list, expression)
	}

	return list
}