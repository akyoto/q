package token

import (
	"github.com/akyoto/q/build/spec"
)

// Tokenize processes the partial read and returns how many bytes were processed.
func Tokenize(buffer []byte, tokens []Token) ([]Token, int) {
	var (
		i              int
		c              byte
		processedBytes int
		lastTokenKind  Kind
		token          = Token{Invalid, nil, 0}
	)

	for i < len(buffer) {
		c = buffer[i]

		switch {
		// Identifiers
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
					i--
					break
				}
			}

			token = Token{Identifier, buffer[processedBytes : i+1], processedBytes}

			if spec.Keywords[string(token.Bytes)] {
				token.Kind = Keyword
			}

		// Numbers
		case (c >= '0' && c <= '9') || (c == '-' && lastTokenKind != Number && lastTokenKind != Identifier && lastTokenKind != GroupEnd && lastTokenKind != ArrayEnd && buffer[i+1] >= '0' && buffer[i+1] <= '9'):
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if c < '0' || c > '9' {
					i--
					break
				}
			}

			token = Token{Number, buffer[processedBytes : i+1], processedBytes}

		// Operators
		case c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>' || c == '!':
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if !(c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>' || c == '!') {
					i--
					break
				}
			}

			token = Token{Operator, buffer[processedBytes : i+1], processedBytes}

			if spec.Operators[string(token.Bytes)] == 0 {
				return tokens, processedBytes
			}

		// Texts
		case c == '"':
			processedBytes = i
			escape := false
			text := make([]byte, 0, 4)

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if escape {
					switch c {
					case 'n':
						text = append(text, '\n')
					case 'r':
						text = append(text, '\r')
					case '\\':
						text = append(text, '\\')
					case '"':
						text = append(text, '"')
					}

					escape = false
					continue
				}

				if c == '\\' {
					escape = true
					continue
				}

				if c == '"' {
					break
				}

				text = append(text, c)
			}

			token = Token{Text, text, processedBytes + 1}

		// Parentheses start
		case c == '(':
			token = Token{GroupStart, nil, i}

		// Parentheses end
		case c == ')':
			token = Token{GroupEnd, nil, i}

		// Block start
		case c == '{':
			token = Token{BlockStart, nil, i}

		// Block end
		case c == '}':
			token = Token{BlockEnd, nil, i}

		// Array start
		case c == '[':
			token = Token{ArrayStart, nil, i}

		// Array end
		case c == ']':
			token = Token{ArrayEnd, nil, i}

		// Separator
		case c == ',':
			token = Token{Separator, nil, i}

		// Accessor
		case c == '.':
			if buffer[i+1] == '.' {
				token = Token{Range, nil, i}
				i++
			} else {
				token = Token{Accessor, nil, i}
			}

		// New line
		case c == '\n':
			token = Token{NewLine, nil, i}
		}

		// Handle token
		if token.Kind != Invalid {
			tokens = append(tokens, token)
			processedBytes = i + 1
			lastTokenKind = token.Kind
			token.Kind = Invalid
		}

		i++
	}

	return tokens, processedBytes
}
