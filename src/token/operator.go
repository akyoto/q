package token

// operator handles all tokens that qualify as an operator.
func operator(tokens List, buffer []byte, i Position) (List, Position) {
	position := i
	kind := Invalid
	increment := false

	for {
		i++
		current := buffer[position:i]

		if i >= Position(len(buffer)) {
			kind, _ = operatorKind(current, 0)
			break
		}

		next := buffer[i]
		kind, increment = operatorKind(current, next)

		if kind != Invalid {
			if increment {
				i++
			}

			break
		}

		if !increment {
			break
		}
	}

	tokens = append(tokens, Token{Kind: kind, Position: position, Length: Length(i - position)})
	return tokens, i
}

func isOperator(c byte) bool {
	switch c {
	case '=', ':', '.', '+', '-', '*', '/', '<', '>', '&', '|', '^', '%', '!':
		return true
	default:
		return false
	}
}