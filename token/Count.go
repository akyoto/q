package token

// Count returns the number of tokens with the same kind.
func Count(tokens []Token, kind Kind) int {
	count := 0

	for _, token := range tokens {
		if token.Kind == kind {
			count++
		}
	}

	return count
}
