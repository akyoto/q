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
			if instruction.Kind != Unknown {
				instruction.Expression = tokens[start:cursor]
				instructions = append(instructions, instruction)
			}

			instruction.Kind = Unknown
			start = cursor + 1

		case token.Operator:
			if instruction.Kind != Unknown {
				continue
			}

			if t.Text() != "=" {
				continue
			}

			instruction.Kind = Assignment

		case token.Keyword:
			if instruction.Kind != Unknown {
				continue
			}

			instruction.Kind = Keyword

		case token.GroupStart:
			if instruction.Kind != Unknown {
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
