package token

import (
	"github.com/akyoto/q/build/spec"
)

// Pre-allocate these byte buffers so we can re-use them
// instead of allocating a new buffer every time.
var (
	groupStartBytes = []byte{'('}
	groupEndBytes   = []byte{')'}
	blockStartBytes = []byte{'{'}
	blockEndBytes   = []byte{'}'}
	arrayStartBytes = []byte{'['}
	arrayEndBytes   = []byte{']'}
	separatorBytes  = []byte{','}
	accessorBytes   = []byte{'.'}
	rangeBytes      = []byte{'.', '.'}
	newLineBytes    = []byte{'\n'}
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

			if spec.Operators[string(token.Bytes)] == nil {
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
			token = Token{GroupStart, groupStartBytes, i}

		// Parentheses end
		case c == ')':
			token = Token{GroupEnd, groupEndBytes, i}

		// Block start
		case c == '{':
			token = Token{BlockStart, blockStartBytes, i}

		// Block end
		case c == '}':
			token = Token{BlockEnd, blockEndBytes, i}

		// Array start
		case c == '[':
			token = Token{ArrayStart, arrayStartBytes, i}

		// Array end
		case c == ']':
			token = Token{ArrayEnd, arrayEndBytes, i}

		// Separator
		case c == ',':
			token = Token{Separator, separatorBytes, i}

		// Accessor
		case c == '.':
			if buffer[i+1] == '.' {
				token = Token{Range, rangeBytes, i}
				i++
			} else {
				token = Token{Accessor, accessorBytes, i}
			}

		// New line
		case c == '\n':
			token = Token{NewLine, newLineBytes, i}
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
