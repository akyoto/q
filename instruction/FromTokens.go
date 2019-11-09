package instruction

import (
	"github.com/akyoto/q/token"
)

// FromTokens takes a list of tokens and generates instructions.
func FromTokens(tokens []token.Token) []Instruction {
	instructions := make([]Instruction, 0, len(tokens)/2)
	start := 0
	instruction := Instruction{}

	for cursor, t := range tokens {
		switch t.Kind {
		case token.NewLine:
			if cursor > start {
				instruction.Tokens = tokens[start:cursor]
				instruction.Position = start
				instructions = append(instructions, instruction)
			}

			instruction.Kind = Invalid
			start = cursor + 1

		case token.Operator:
			if instruction.Kind != Invalid {
				continue
			}

			if t.Text() != "=" {
				continue
			}

			instruction.Kind = Assignment

		case token.Keyword:
			if instruction.Kind != Invalid {
				continue
			}

			instruction.Kind = Keyword

		case token.GroupStart:
			if instruction.Kind != Invalid {
				continue
			}

			if tokens[cursor-1].Kind != token.Identifier {
				continue
			}

			instruction.Kind = Call
		}
	}

	return instructions
}
