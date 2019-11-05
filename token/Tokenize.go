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
		token          = Token{Unknown, ""}
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

			token = Token{Identifier, string(buffer[processedBytes : i+1])}

			if spec.Keywords[token.Text] {
				token.Kind = Keyword
			}

		// Numbers
		case c >= '0' && c <= '9':
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]

				if c < '0' || c > '9' {
					i--
					break
				}
			}

			token = Token{Number, string(buffer[processedBytes : i+1])}

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

			token = Token{Text, string(buffer[processedBytes : i+1])}

		// Parentheses start
		case c == '(':
			token = Token{GroupStart, string(buffer[i : i+1])}

		// Parentheses end
		case c == ')':
			token = Token{GroupEnd, string(buffer[i : i+1])}

		// Block start
		case c == '{':
			token = Token{BlockStart, string(buffer[i : i+1])}

		// Block end
		case c == '}':
			token = Token{BlockEnd, string(buffer[i : i+1])}

		// New line
		case c == '\n':
			token = Token{NewLine, ""}

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

			token = Token{WhiteSpace, string(buffer[processedBytes : i+1])}
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
