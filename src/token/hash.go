package token

// hash handles all tokens starting with '#'.
func hash(tokens List, buffer []byte, i Position) (List, Position) {
	if i+1 < Position(len(buffer)) && buffer[i+1] == '!' {
		position := i

		for i < Position(len(buffer)) && buffer[i] != '\n' {
			i++
		}

		tokens = append(tokens, Token{Kind: Script, Position: position, Length: Length(i - position)})
	} else {
		tokens = append(tokens, Token{Kind: Invalid, Position: i, Length: 1})
		i++
	}

	return tokens, i
}