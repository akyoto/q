package token

// Index returns the position of a token kind within a token list.
func Index(tokens []Token, kind Kind) Position {
	for i, token := range tokens {
		if token.Kind == kind {
			return i
		}
	}

	return -1
}
