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
			{token.Identifier, 0, []byte("abc")},
			{token.GroupStart, 3, []byte{'('}},
			{token.GroupEnd, 4, []byte{')'}},
			{token.NewLine, 5, []byte{'\n'}},
		}},
		{[]byte("x = 5\n"), []token.Token{
			{token.Identifier, 0, []byte("x")},
			{token.Operator, 2, []byte("=")},
			{token.Number, 4, []byte("5")},
			{token.NewLine, 5, []byte{'\n'}},
		}},
		{[]byte("for i = 0..2\n"), []token.Token{
			{token.Keyword, 0, []byte("for")},
			{token.Identifier, 4, []byte("i")},
			{token.Operator, 6, []byte("=")},
			{token.Number, 8, []byte("0")},
			{token.Range, 9, []byte("..")},
			{token.Number, 11, []byte("2")},
			{token.NewLine, 12, []byte{'\n'}},
		}},
		{[]byte("abc() {\n 123 = \"text\", return}\n"), []token.Token{
			{token.Identifier, 0, []byte("abc")},
			{token.GroupStart, 3, []byte{'('}},
			{token.GroupEnd, 4, []byte{')'}},
			{token.BlockStart, 6, []byte{'{'}},
			{token.NewLine, 7, []byte{'\n'}},
			{token.Number, 9, []byte("123")},
			{token.Operator, 13, []byte("=")},
			{token.Text, 16, []byte("text")},
			{token.Separator, 21, []byte{','}},
			{token.Keyword, 23, []byte("return")},
			{token.BlockEnd, 29, []byte{'}'}},
			{token.NewLine, 30, []byte{'\n'}},
		}},
		{[]byte("# A comment.\n"), []token.Token{
			{token.Comment, 0, []byte("A comment.")},
			{token.NewLine, 12, []byte{'\n'}},
		}},
	}

	for _, pattern := range usagePatterns {
		tokens := []token.Token{}
		processed := uint16(0)
		tokens, processed = token.Tokenize(pattern.Source, tokens)
		assert.Equal(t, processed, uint16(len(pattern.Source)))

		for index := range tokens {
			assert.Equal(t, tokens[index].Kind, pattern.Expected[index].Kind)
			assert.DeepEqual(t, tokens[index].Bytes, pattern.Expected[index].Bytes)
			assert.Equal(t, tokens[index].Position, pattern.Expected[index].Position)
			assert.Equal(t, tokens[index].Text(), pattern.Expected[index].Text())
			assert.Equal(t, tokens[index].Kind.String(), pattern.Expected[index].Kind.String())
		}
	}
}
