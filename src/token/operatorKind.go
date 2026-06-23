package token

// operatorKind returns the token kind for the given byte slice.
// The boolean value indicates if the next byte was used.
// A valid token always stops the search for more characters.
// An invalid token with true indicates continued parsing.
// An invalid token with false indicates cancellation.
func operatorKind(op []byte, next byte) (Kind, bool) {
	switch string(op) {
	case "!":
		switch next {
		case '=':
			return NotEqual, true
		default:
			return Not, false
		}
	case "%":
		switch next {
		case '=':
			return ModAssign, true
		default:
			return Mod, false
		}
	case "&":
		switch next {
		case '=':
			return AndAssign, true
		case '&':
			return LogicalAnd, true
		default:
			return And, false
		}
	case "*":
		switch next {
		case '=':
			return MulAssign, true
		default:
			return Mul, false
		}
	case "+":
		switch next {
		case '=':
			return AddAssign, true
		default:
			return Add, false
		}
	case ".":
		switch next {
		case '.':
			return Range, true
		default:
			return Dot, false
		}
	case ":":
		switch next {
		case '=':
			return Define, true
		default:
			return FieldAssign, false
		}
	case "<":
		switch next {
		case '=':
			return LessEqual, true
		case '<':
			return Invalid, true
		default:
			return Less, false
		}
	case "<<":
		switch next {
		case '=':
			return ShlAssign, true
		default:
			return Shl, false
		}
	case ">":
		switch next {
		case '=':
			return GreaterEqual, true
		case '>':
			return Invalid, true
		default:
			return Greater, false
		}
	case ">>":
		switch next {
		case '=':
			return ShrAssign, true
		default:
			return Shr, false
		}
	case "=":
		switch next {
		case '=':
			return Equal, true
		default:
			return Assign, false
		}
	case "^":
		switch next {
		case '=':
			return XorAssign, true
		default:
			return Xor, false
		}
	case "|":
		switch next {
		case '=':
			return OrAssign, true
		case '|':
			return LogicalOr, true
		default:
			return Or, false
		}
	default:
		return Invalid, false
	}
}