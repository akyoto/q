package expression

import "git.urbach.dev/cli/q/src/token"

// newTypeExpression creates a new type expression.
func newTypeExpression(tokens token.List) *TypeExpression {
	start := tokens[0].Position
	end := token.Position(tokens[len(tokens)-1].End())

	virtualToken := token.Token{
		Position: start,
		Length:   token.Length(end - start),
		Kind:     token.Identifier,
	}

	return &TypeExpression{
		Tokens:     tokens,
		Expression: Expression{Token: virtualToken},
	}
}