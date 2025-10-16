package token

// digit handles all tokens that qualify as a decimal digit.
func digit(tokens List, buffer []byte, i Position) (List, Position) {
	position := i
	i++

	for i < Position(len(buffer)) && isDigit(buffer[i]) {
		i++
	}

	last := len(tokens) - 1

	if len(tokens) > 0 && tokens[last].Kind == Negate {
		tokens[last].Kind = Number
		tokens[last].Length = Length(i-position) + 1
	} else {
		tokens = append(tokens, Token{Kind: Number, Position: position, Length: Length(i - position)})
	}

	return tokens, i
}

// isDigit returns true for decimal digits.
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}