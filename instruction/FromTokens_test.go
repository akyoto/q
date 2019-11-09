package instruction_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/instruction"
	"github.com/akyoto/q/token"
)

func TestFromTokens(t *testing.T) {
	source := []byte("\na = 1\nb()\nreturn\n")
	tokens := []token.Token{}
	processed := 0
	tokens, processed = token.Tokenize(source, tokens)
	assert.Equal(t, processed, len(source))
	assert.Equal(t, len(tokens), 11)
	instructions := instruction.FromTokens(tokens)
	assert.Equal(t, len(instructions), 3)
	assert.Equal(t, instructions[0].Kind, instruction.Assignment)
	assert.Equal(t, instructions[1].Kind, instruction.Call)
	assert.Equal(t, instructions[2].Kind, instruction.Keyword)
}
