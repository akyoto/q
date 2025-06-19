package token

// Tokenize turns the file contents into a list of tokens.
func Tokenize(buffer []byte) List {
	var (
		i      Position
		tokens = make(List, 0, 8+len(buffer)/2)
	)

	for i < Position(len(buffer)) {
		switch buffer[i] {
		case ' ', '\t', '\r':
		case ',':
			tokens = append(tokens, Token{Kind: Separator, Position: i, Length: 1})
		case '(':
			tokens = append(tokens, Token{Kind: GroupStart, Position: i, Length: 1})
		case ')':
			tokens = append(tokens, Token{Kind: GroupEnd, Position: i, Length: 1})
		case '{':
			tokens = append(tokens, Token{Kind: BlockStart, Position: i, Length: 1})
		case '}':
			tokens = append(tokens, Token{Kind: BlockEnd, Position: i, Length: 1})
		case '[':
			tokens = append(tokens, Token{Kind: ArrayStart, Position: i, Length: 1})
		case ']':
			tokens = append(tokens, Token{Kind: ArrayEnd, Position: i, Length: 1})
		case '\n':
			tokens = append(tokens, Token{Kind: NewLine, Position: i, Length: 1})
		case '-':
			tokens, i = dash(tokens, buffer, i)
		case '/':
			tokens, i = slash(tokens, buffer, i)
			continue
		case '"', '\'':
			tokens, i = quote(tokens, buffer, i)
			continue
		case '0':
			tokens, i = zero(tokens, buffer, i)
			continue
		default:
			if isIdentifierStart(buffer[i]) {
				tokens, i = identifier(tokens, buffer, i)
				continue
			}

			if isDigit(buffer[i]) {
				tokens, i = digit(tokens, buffer, i)
				continue
			}

			if isOperator(buffer[i]) {
				tokens, i = operator(tokens, buffer, i)
				continue
			}

			tokens = append(tokens, Token{Kind: Invalid, Position: i, Length: 1})
		}

		i++
	}

	tokens = append(tokens, Token{Kind: EOF, Position: i, Length: 0})
	return tokens
}