package instruction

import (
	"fmt"

	"github.com/akyoto/q/build/token"
)

// FromTokens takes a list of tokens and generates instructions.
func FromTokens(tokens []token.Token) ([]Instruction, *Error) {
	instructions := make([]Instruction, 0, len(tokens)/2)
	start := 0
	instruction := Instruction{}
	groups := 0
	blocks := []Kind{}

	for i, t := range tokens {
		switch t.Kind {
		case token.NewLine:
			if start == i {
				start = i + 1
				continue
			}

			if instruction.Kind != Assignment && instruction.Kind != Return && instruction.Kind != Invalid {
				continue
			}

			instruction.Tokens = tokens[start:i]
			instruction.Position = start
			instructions = append(instructions, instruction)

			instruction.Kind = Invalid
			start = i + 1

		case token.Operator:
			if instruction.Kind != Invalid {
				continue
			}

			if t.Text() != "=" {
				continue
			}

			instruction.Kind = Assignment

		case token.GroupStart:
			groups++

			if groups != 1 {
				continue
			}

			if instruction.Kind != Invalid {
				continue
			}

			if tokens[i-1].Kind != token.Identifier {
				continue
			}

			instruction.Kind = Call

		case token.GroupEnd:
			groups--

			if groups != 0 {
				continue
			}

			if instruction.Kind != Call {
				continue
			}

			instruction.Tokens = tokens[start : i+1]
			instruction.Position = start
			instructions = append(instructions, instruction)

			instruction.Kind = Invalid
			start = i + 1

		case token.Identifier:
			if instruction.Kind != Invalid {
				continue
			}

			if i == len(tokens)-1 {
				return instructions, &Error{fmt.Sprintf("Expected assignment or function call after '%s'", t.Text()), i, true}
			}

			nextToken := tokens[i+1]

			if nextToken.Kind != token.Operator && nextToken.Kind != token.GroupStart {
				remaining := tokens[i+1:]
				newLinePos := token.Index(remaining, token.NewLine)

				if newLinePos != -1 && remaining[newLinePos-1].Kind == token.GroupEnd {
					return instructions, &Error{fmt.Sprintf("Missing opening bracket '(' after '%s'", t.Text()), i, true}
				}

				return instructions, &Error{fmt.Sprintf("Expected assignment or function call after '%s'", t.Text()), i, true}
			}

		case token.Keyword:
			if instruction.Kind != Invalid {
				continue
			}

			switch t.Text() {
			case "if":
				instruction.Kind = IfStart
			case "for":
				instruction.Kind = ForStart
			case "loop":
				instruction.Kind = LoopStart
			case "return":
				instruction.Kind = Return
			default:
				return instructions, &Error{"Keyword not implemented", i, false}
			}

		case token.BlockStart:
			switch instruction.Kind {
			case IfStart, ForStart, LoopStart:
				// OK.

			default:
				return instructions, &Error{fmt.Sprintf("Invalid block of type '%s'", instruction.Kind), i, false}
			}

			blocks = append(blocks, instruction.Kind)

			instruction.Tokens = tokens[start:i]
			instruction.Position = start
			instructions = append(instructions, instruction)

			instruction.Kind = Invalid
			start = i + 1

		case token.BlockEnd:
			block := blocks[len(blocks)-1]

			switch block {
			case IfStart:
				instruction.Kind = IfEnd

			case ForStart:
				instruction.Kind = ForEnd

			case LoopStart:
				instruction.Kind = LoopEnd

			default:
				return instructions, &Error{fmt.Sprintf("Not implemented: %v", block), i, false}
			}

			instruction.Tokens = tokens[start:i]
			instruction.Position = start
			instructions = append(instructions, instruction)

			instruction.Kind = Invalid
			start = i + 1

			blocks = blocks[:len(blocks)-1]
		}
	}

	if start != len(tokens) {
		instruction.Tokens = tokens[start:]
		instruction.Position = start
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}
