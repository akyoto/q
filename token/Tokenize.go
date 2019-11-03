package token

import "github.com/akyoto/q/spec"

// Tokenize processes the partial read and returns how many bytes were processed.
// The remaining bytes will be prepended to the next call.
// The custom function handleToken is called for each processed token.
func Tokenize(buffer []byte, handleToken func(Token) error) (int, error) {
	var (
		i              int
		c              byte
		processedBytes int
		token          = Token{Unknown, nil}
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
					return processedBytes, nil
				}

				c = buffer[i]

				if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
					i--
					break
				}
			}

			token = Token{Identifier, buffer[processedBytes : i+1]}

			if spec.Keywords[string(token.Text)] {
				token.Kind = Keyword
			}

		// Texts
		case c == '"':
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]

				if c == '"' {
					break
				}
			}

			token = Token{Text, buffer[processedBytes : i+1]}

		// Parentheses start
		case c == '(':
			token = Token{GroupStart, buffer[i : i+1]}

		// Parentheses end
		case c == ')':
			token = Token{GroupEnd, buffer[i : i+1]}

		// Block start
		case c == '{':
			token = Token{BlockStart, buffer[i : i+1]}

		// Block end
		case c == '}':
			token = Token{BlockEnd, buffer[i : i+1]}

		// End of line
		case c == '\n':
			token = Token{StartOfLine, nil}

		// Whitespace
		case c == ' ' || c == '\t':
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]

				if c != ' ' && c != '\t' {
					i--
					break
				}
			}

			token = Token{WhiteSpace, buffer[processedBytes : i+1]}
		}

		// Handle token
		if token.Kind != Unknown {
			err := handleToken(token)

			if err != nil {
				return 0, err
			}

			processedBytes = i + 1
			token.Kind = Unknown
		}

		i++
	}

	return processedBytes, nil
}
