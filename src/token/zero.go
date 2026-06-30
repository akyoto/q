package token

// zero handles all tokens starting with a '0'.
func zero(tokens List, buffer []byte, i Position) (List, Position) {
	position := i
	i++

	if i >= Position(len(buffer)) {
		tokens = append(tokens, Token{Kind: Number, Position: position, Length: 1})
		return tokens, i
	}

	filter := isDigit

	switch buffer[i] {
	case 'x':
		i++
		filter = isHexDigit

	case 'b':
		i++
		filter = isBinaryDigit

	case 'o':
		i++
		filter = isOctalDigit
	}

	for i < Position(len(buffer)) && filter(buffer[i]) {
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

// isHexDigit returns true if the character represents a hexadecimal digit.
func isHexDigit(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'F')
}

// isBinaryDigit returns true if the character represents a binary digit.
func isBinaryDigit(c byte) bool {
	return c == '0' || c == '1'
}

// isOctalDigit returns true if the character represents an octal digit.
func isOctalDigit(c byte) bool {
	return c >= '0' && c <= '7'
}