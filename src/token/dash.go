package token

// dash handles all tokens starting with '-'.
func dash(tokens List, buffer []byte, i Position) (List, Position) {
	if len(tokens) == 0 || tokens[len(tokens)-1].IsOperator() || tokens[len(tokens)-1].IsExpressionStart() || tokens[len(tokens)-1].IsKeyword() {
		tokens = append(tokens, Token{Kind: Negate, Position: i, Length: 1})
	} else {
		if i+1 < Position(len(buffer)) {
			switch buffer[i+1] {
			case '=':
				tokens = append(tokens, Token{Kind: SubAssign, Position: i, Length: 2})
				i++
			case '>':
				tokens = append(tokens, Token{Kind: ReturnType, Position: i, Length: 2})
				i++
			default:
				tokens = append(tokens, Token{Kind: Sub, Position: i, Length: 1})
			}
		} else {
			tokens = append(tokens, Token{Kind: Sub, Position: i, Length: 1})
		}
	}

	return tokens, i
}