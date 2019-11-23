package instruction_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/token"
)

func TestFromTokens(t *testing.T) {
	usagePatterns := []struct {
		Source   []byte
		Expected []instruction.Instruction
	}{
		{[]byte("\na = 1\nb()\nloop\n{\na)\n}\n"), []instruction.Instruction{
			{instruction.Assignment, nil, 1},
			{instruction.Call, nil, 5},
			{instruction.LoopStart, nil, 9},
			{instruction.Invalid, nil, 13},
			{instruction.LoopEnd, nil, 16},
		}},
		{[]byte("if x > 1 {\nx = 2\n}\n"), []instruction.Instruction{
			{instruction.IfStart, nil, 0},
			{instruction.Assignment, nil, 6},
			{instruction.IfEnd, nil, 10},
		}},
		{[]byte("for i = 0..2 {}\n"), []instruction.Instruction{
			{instruction.ForStart, nil, 0},
			{instruction.ForEnd, nil, 7},
		}},
	}

	for _, pattern := range usagePatterns {
		tokens := []token.Token{}
		processed := 0
		tokens, processed = token.Tokenize(pattern.Source, tokens)
		assert.Equal(t, processed, len(pattern.Source))
		instructions := instruction.FromTokens(tokens)
		assert.Equal(t, len(instructions), len(pattern.Expected))

		for index := range instructions {
			t.Logf("[%d][%s] %s", index, instructions[index].Kind, instructions[index].Tokens)
			assert.Equal(t, instructions[index].Kind, pattern.Expected[index].Kind)
			assert.Equal(t, instructions[index].Position, pattern.Expected[index].Position)
			assert.Equal(t, instructions[index].Kind.String(), pattern.Expected[index].Kind.String())
		}
	}
}
