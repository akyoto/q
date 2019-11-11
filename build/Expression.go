package build

import "github.com/akyoto/q/token"

// Expression is a list of tokens.
type Expression []token.Token

// First returns the first token in the expression.
func (e Expression) First() token.Token {
	return e[0]
}

// Last returns the last token in the expression.
func (e Expression) Last() token.Token {
	return e[len(e)-1]
}
