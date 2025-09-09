package token

// identifier handles all tokens that qualify as an identifier.
func identifier(tokens List, buffer []byte, i Position) (List, Position) {
	position := i
	i++

	for i < Position(len(buffer)) && isIdentifier(buffer[i]) {
		i++
	}

	identifier := buffer[position:i]
	kind := Identifier

	switch string(identifier) {
	case "as":
		kind = Cast
	case "assert":
		kind = Assert
	case "const":
		kind = Const
	case "if":
		kind = If
	case "else":
		kind = Else
	case "extern":
		kind = Extern
	case "import":
		kind = Import
	case "loop":
		kind = Loop
	case "return":
		kind = Return
	case "switch":
		kind = Switch
	}

	tokens = append(tokens, Token{Kind: kind, Position: position, Length: Length(len(identifier))})
	return tokens, i
}

func isIdentifier(c byte) bool {
	return isLetter(c) || isDigit(c) || c == '_'
}

func isIdentifierStart(c byte) bool {
	return isLetter(c) || c == '_'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}