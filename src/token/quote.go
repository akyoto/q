package token

// quote handles all tokens starting with a single or double quote.
func quote(tokens List, buffer []byte, i Position) (List, Position) {
	limiter := buffer[i]
	start := i
	end := Position(len(buffer))
	i++

	for i < Position(len(buffer)) {
		if buffer[i] == limiter && (buffer[i-1] != '\\' || buffer[i-2] == '\\') {
			end = i + 1
			i++
			break
		}

		i++
	}

	kind := String

	if limiter == '\'' {
		kind = Rune
	}

	tokens = append(tokens, Token{Kind: kind, Position: start, Length: Length(end - start)})
	return tokens, i
}