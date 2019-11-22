package instruction_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/token"
)

func TestFromTokens(t *testing.T) {
	source := []byte("\na = 1\nb()\nloop\n{\na)\n}\n")
	expected := []instruction.Instruction{
		{instruction.Assignment, nil, 1},
		{instruction.Call, nil, 5},
		{instruction.LoopStart, nil, 9},
		{instruction.Invalid, nil, 13},
		{instruction.LoopEnd, nil, 16},
	}

	tokens := []token.Token{}
	processed := 0
	tokens, processed = token.Tokenize(source, tokens)
	assert.Equal(t, processed, len(source))
	assert.Equal(t, len(tokens), 18)
	instructions := instruction.FromTokens(tokens)
	assert.Equal(t, len(instructions), len(expected))

	for index := range instructions {
		t.Logf("[%d][%s] %s", index, instructions[index].Kind, instructions[index].Tokens)
		assert.Equal(t, instructions[index].Kind, expected[index].Kind)
		assert.Equal(t, instructions[index].Position, expected[index].Position)
		assert.Equal(t, instructions[index].Kind.String(), expected[index].Kind.String())
	}
}

func TestFromTokensIf(t *testing.T) {
	source := []byte("if x > 1 {\nx = 2\n}\n")
	expected := []instruction.Instruction{
		{instruction.IfStart, nil, 0},
		{instruction.Assignment, nil, 6},
		{instruction.IfEnd, nil, 10},
	}

	tokens := []token.Token{}
	processed := 0
	tokens, processed = token.Tokenize(source, tokens)
	assert.Equal(t, processed, len(source))
	assert.Equal(t, len(tokens), 12)
	instructions := instruction.FromTokens(tokens)
	assert.Equal(t, len(instructions), len(expected))

	for index := range instructions {
		t.Logf("[%d][%s] %s", index, instructions[index].Kind, instructions[index].Tokens)
		assert.Equal(t, instructions[index].Kind, expected[index].Kind)
		assert.Equal(t, instructions[index].Position, expected[index].Position)
		assert.Equal(t, instructions[index].Kind.String(), expected[index].Kind.String())
	}
}
