package expression

import "git.urbach.dev/cli/q/src/token"

// TypeExpression is an expression that saves the tokens for data type parsing later.
type TypeExpression struct {
	Expression
	Tokens []token.Token
}