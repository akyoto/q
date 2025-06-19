package token

// operator handles all tokens that qualify as an operator.
func operator(tokens List, buffer []byte, i Position) (List, Position) {
	position := i
	i++

	for i < Position(len(buffer)) && isOperator(buffer[i]) {
		i++
	}

	kind := Invalid

	switch string(buffer[position:i]) {
	case "!":
		kind = Not
	case "!=":
		kind = NotEqual
	case "%":
		kind = Mod
	case "%=":
		kind = ModAssign
	case "&":
		kind = And
	case "&&":
		kind = LogicalAnd
	case "&=":
		kind = AndAssign
	case "*":
		kind = Mul
	case "*=":
		kind = MulAssign
	case "+":
		kind = Add
	case "+=":
		kind = AddAssign
	case ".":
		kind = Dot
	case "..":
		kind = Range
	case ":=":
		kind = Define
	case "<":
		kind = Less
	case "<<":
		kind = Shl
	case "<<=":
		kind = ShlAssign
	case "<=":
		kind = LessEqual
	case "=":
		kind = Assign
	case "==":
		kind = Equal
	case ">":
		kind = Greater
	case ">=":
		kind = GreaterEqual
	case ">>":
		kind = Shr
	case ">>=":
		kind = ShrAssign
	case "^":
		kind = Xor
	case "^=":
		kind = XorAssign
	case "|":
		kind = Or
	case "|=":
		kind = OrAssign
	case "||":
		kind = LogicalOr
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