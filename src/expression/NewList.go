package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// NewList generates a list of expressions from comma separated parameters.
func NewList(tokens token.List) []*Expression {
	var list []*Expression

	for param := range tokens.Split {
		expression := Parse(param)
		list = append(list, expression)
	}

	return list
}