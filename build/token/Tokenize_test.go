package token_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/token"
)

func TestTokenize(t *testing.T) {
	usagePatterns := []struct {
		Source   []byte
		Expected []token.Token
	}{
		{[]byte("abc()\n"), []token.Token{
			{token.Identifier, []byte("abc"), 0},
			{token.GroupStart, []byte{'('}, 3},
			{token.GroupEnd, []byte{')'}, 4},
			{token.NewLine, []byte{'\n'}, 5},
		}},
		{[]byte("x = 5\n"), []token.Token{
			{token.Identifier, []byte("x"), 0},
			{token.Operator, []byte("="), 2},
			{token.Number, []byte("5"), 4},
			{token.NewLine, []byte{'\n'}, 5},
		}},
		{[]byte("for i = 0..2\n"), []token.Token{
			{token.Keyword, []byte("for"), 0},
			{token.Identifier, []byte("i"), 4},
			{token.Operator, []byte("="), 6},
			{token.Number, []byte("0"), 8},
			{token.Range, []byte(".."), 9},
			{token.Number, []byte("2"), 11},
			{token.NewLine, []byte{'\n'}, 12},
		}},
		{[]byte("abc() {\n 123 = \"text\", return}\n"), []token.Token{
			{token.Identifier, []byte("abc"), 0},
			{token.GroupStart, []byte{'('}, 3},
			{token.GroupEnd, []byte{')'}, 4},
			{token.BlockStart, []byte{'{'}, 6},
			{token.NewLine, []byte{'\n'}, 7},
			{token.Number, []byte("123"), 9},
			{token.Operator, []byte("="), 13},
			{token.Text, []byte("text"), 16},
			{token.Separator, []byte{','}, 21},
			{token.Keyword, []byte("return"), 23},
			{token.BlockEnd, []byte{'}'}, 29},
			{token.NewLine, []byte{'\n'}, 30},
		}},
	}

	for _, pattern := range usagePatterns {
		tokens := []token.Token{}
		processed := 0
		tokens, processed = token.Tokenize(pattern.Source, tokens)
		assert.Equal(t, processed, len(pattern.Source))
		assert.DeepEqual(t, tokens, pattern.Expected)

		for index := range tokens {
			assert.Equal(t, tokens[index].Kind.String(), pattern.Expected[index].Kind.String())
		}
	}
}
