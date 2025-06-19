package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestCount(t *testing.T) {
	buffer := []byte(`a b b c c c`)
	tokens := token.Tokenize(buffer)
	assert.Equal(t, token.Count(tokens, buffer, token.Identifier, "a"), 1)
	assert.Equal(t, token.Count(tokens, buffer, token.Identifier, "b"), 2)
	assert.Equal(t, token.Count(tokens, buffer, token.Identifier, "c"), 3)
	assert.Equal(t, token.Count(tokens, buffer, token.Identifier, "d"), 0)
}