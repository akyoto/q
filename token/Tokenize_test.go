package token_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/token"
)

func TestTokenize(t *testing.T) {
	source := []byte("func abc() {\n 123 \"text\" }")
	expected := []token.Token{
		{token.Keyword, []byte("func"), 0},
		{token.Identifier, []byte("abc"), 5},
		{token.GroupStart, nil, 8},
		{token.GroupEnd, nil, 9},
		{token.BlockStart, nil, 11},
		{token.NewLine, nil, 12},
		{token.Number, []byte("123"), 14},
		{token.Text, []byte("text"), 19},
		{token.BlockEnd, nil, 25},
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
