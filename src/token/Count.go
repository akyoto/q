package token

// Count counts how often the given token appears in the token list.
func Count(tokens []Token, buffer []byte, kind Kind, name string) uint8 {
	count := uint8(0)

	for _, t := range tokens {
		if t.Kind == kind && t.String(buffer) == name {
			count++
		}
	}

	return count
}