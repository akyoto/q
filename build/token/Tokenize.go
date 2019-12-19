package token

import (
	"bytes"

	"github.com/akyoto/q/build/keywords"
	"github.com/akyoto/q/build/operators"
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
	questionBytes   = []byte{'?'}
	newLineBytes    = []byte{'\n'}
)

// Tokenize processes the partial read and returns how many bytes were processed.
func Tokenize(buffer []byte, tokens []Token) ([]Token, uint16) {
	var (
		i              uint16
		c              byte
		processedBytes uint16
		lastTokenKind  Kind
		token          = Token{Invalid, 0, nil}
	)

	for i < uint16(len(buffer)) {
		c = buffer[i]

		switch {
		// Identifiers
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_':
			processedBytes = i

			for {
				i++

				if i >= uint16(len(buffer)) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') && c != '_' {
					i--
					break
				}
			}

			token = Token{Identifier, processedBytes, buffer[processedBytes : i+1]}

			if keywords.All[string(token.Bytes)] {
				token.Kind = Keyword
			}

		// Numbers
		case (c >= '0' && c <= '9') || (c == '-' && lastTokenKind != Number && lastTokenKind != Identifier && lastTokenKind != GroupEnd && lastTokenKind != ArrayEnd && buffer[i+1] >= '0' && buffer[i+1] <= '9'):
			processedBytes = i

			for {
				i++

				if i >= uint16(len(buffer)) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if c < '0' || c > '9' {
					i--
					break
				}
			}

			token = Token{Number, processedBytes, buffer[processedBytes : i+1]}

		case c == '#':
			processedBytes = i

			for {
				i++

				if i >= uint16(len(buffer)) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if c == '\n' {
					i--
					break
				}
			}

			trimmed := bytes.TrimSpace(buffer[processedBytes+1 : i+1])
			token = Token{Comment, processedBytes, trimmed}

		// Operators
		case c == '=' || c == ':' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>' || c == '!':
			processedBytes = i

			for {
				i++

				if i >= uint16(len(buffer)) {
					return tokens, processedBytes
				}

				c = buffer[i]

				if !(c == '=' || c == ':' || c == '+' || c == '-' || c == '*' || c == '/' || c == '<' || c == '>' || c == '!') {
					i--
					break
				}
			}

			token = Token{Operator, processedBytes, buffer[processedBytes : i+1]}

			if operators.All[string(token.Bytes)] == nil {
				return tokens, processedBytes
			}

		// Texts
		case c == '"':
			processedBytes = i
			escape := false
			text := make([]byte, 0, 4)

			for {
				i++

				if i >= uint16(len(buffer)) {
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

			token = Token{Text, processedBytes + 1, text}

		// Parentheses start
		case c == '(':
			token = Token{GroupStart, i, groupStartBytes}

		// Parentheses end
		case c == ')':
			token = Token{GroupEnd, i, groupEndBytes}

		// Block start
		case c == '{':
			token = Token{BlockStart, i, blockStartBytes}

		// Block end
		case c == '}':
			token = Token{BlockEnd, i, blockEndBytes}

		// Array start
		case c == '[':
			token = Token{ArrayStart, i, arrayStartBytes}

		// Array end
		case c == ']':
			token = Token{ArrayEnd, i, arrayEndBytes}

		// Separator
		case c == ',':
			token = Token{Separator, i, separatorBytes}

		// Accessor
		case c == '.':
			if buffer[i+1] == '.' {
				token = Token{Range, i, rangeBytes}
				i++
			} else {
				token = Token{Operator, i, accessorBytes}
			}

		// Question
		case c == '?':
			token = Token{Question, i, questionBytes}

		// New line
		case c == '\n':
			token = Token{NewLine, i, newLineBytes}
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
