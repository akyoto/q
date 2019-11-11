package expression

import (
	"errors"

	"github.com/akyoto/q/build/token"
)

// Expression is a binary tree with an operator on each node.
type Expression struct {
	Value    token.Token
	Operands []*Expression
}

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	for _, t := range tokens {
		if t.Kind == token.Identifier || t.Kind == token.Text || t.Kind == token.Number {
			return &Expression{Value: t}, nil
		}
	}

	return nil, errors.New("Invalid expression")
}
