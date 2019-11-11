package expression_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

func TestExpression(t *testing.T) {
	src := []byte("3\n")
	tokens, processed := token.Tokenize(src, []token.Token{})
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, tokens[0].Kind, token.Number)
	assert.Equal(t, processed, len(src))
	expr, err := expression.FromTokens(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, expr)
}
