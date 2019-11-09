package instruction_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/instruction"
	"github.com/akyoto/q/token"
)

func TestFromTokens(t *testing.T) {
	source := []byte("\na = 1\nb()\nreturn\na)\n")
	expected := []instruction.Instruction{
		{instruction.Assignment, nil, 1},
		{instruction.Call, nil, 5},
		{instruction.Keyword, nil, 9},
		{instruction.Invalid, nil, 11},
	}

	tokens := []token.Token{}
	processed := 0
	tokens, processed = token.Tokenize(source, tokens)
	assert.Equal(t, processed, len(source))
	assert.Equal(t, len(tokens), 14)
	instructions := instruction.FromTokens(tokens)
	assert.Equal(t, len(instructions), 4)

	for index := range instructions {
		assert.Equal(t, instructions[index].Kind, expected[index].Kind)
		assert.Equal(t, instructions[index].Position, expected[index].Position)
		assert.Equal(t, instructions[index].Kind.String(), expected[index].Kind.String())
	}
}
