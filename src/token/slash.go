package token

// slash handles all tokens starting with '/'.
func slash(tokens List, buffer []byte, i Position) (List, Position) {
	if i+1 < Position(len(buffer)) && buffer[i+1] == '/' {
		position := i

		for i < Position(len(buffer)) && buffer[i] != '\n' {
			i++
		}

		tokens = append(tokens, Token{Kind: Comment, Position: position, Length: Length(i - position)})
	} else {
		position := i
		i++

		for i < Position(len(buffer)) && isOperator(buffer[i]) {
			i++
		}

		kind := Invalid

		switch string(buffer[position:i]) {
		case "/":
			kind = Div
		case "/=":
			kind = DivAssign
		}

		tokens = append(tokens, Token{Kind: kind, Position: position, Length: Length(i - position)})
	}

	return tokens, i
}