package token

// Index returns the position of a token within a token list.
func Index(tokens []Token, kind Kind, text string) Position {
	for i, token := range tokens {
		if token.Kind == kind && token.Text() == text {
			return i
		}
	}

	return -1
}

// IndexKind returns the position of a token kind within a token list.
func IndexKind(tokens []Token, kind Kind) Position {
	for i, token := range tokens {
		if token.Kind == kind {
			return i
		}
	}

	return -1
}
