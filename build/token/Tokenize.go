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
		case c >= '0' && c <= '9':
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
		case c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>':
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if !(c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>') {
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

			for {
				i++

				if i >= len(buffer) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if c == '"' {
					break
				}
			}

			token = Token{Text, buffer[processedBytes+1 : i], processedBytes + 1}

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

		// Separator
		case c == ',':
			token = Token{Separator, nil, i}

		// New line
		case c == '\n':
			token = Token{NewLine, nil, i}
		}

		// Handle token
		if token.Kind != Invalid {
			tokens = append(tokens, token)
			processedBytes = i + 1
			token.Kind = Invalid
		}

		i++
	}

	return tokens, processedBytes
}
