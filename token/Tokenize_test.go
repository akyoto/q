package token_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/token"
)

func TestTokenize(t *testing.T) {
	source := []byte("abc() {\n 123 \"text\" }")
	expected := []token.Token{
		{token.Identifier, []byte("abc"), 0},
		{token.GroupStart, nil, 3},
		{token.GroupEnd, nil, 4},
		{token.BlockStart, nil, 6},
		{token.NewLine, nil, 7},
		{token.Number, []byte("123"), 9},
		{token.Text, []byte("text"), 14},
		{token.BlockEnd, nil, 20},
	}

	tokens := []token.Token{}
	processed := 0
	tokens, processed = token.Tokenize(source, tokens)
	assert.Equal(t, processed, len(source))
	assert.DeepEqual(t, tokens, expected)

	for index := range tokens {
		assert.Equal(t, tokens[index].Kind.String(), expected[index].Kind.String())
	}
}
