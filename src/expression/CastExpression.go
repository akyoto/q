package expression

import "git.urbach.dev/cli/q/src/token"

// CastExpression is an expression that saves the tokens for data type parsing later.
type CastExpression struct {
	Expression
	Tokens []token.Token
}