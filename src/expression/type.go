package expression

import "git.urbach.dev/cli/q/src/token"

// allowedInType returns true if the given `kind` can be in a type expression.
func allowedInType(kind token.Kind) bool {
	switch kind {
	case token.Identifier, token.Dot, token.Mul, token.Not, token.ArrayStart, token.ArrayEnd:
		return true
	}

	return false
}

// makeTypeToken creates a new type expression wrapped inside a single token.
func makeTypeToken(tokens token.List) token.Token {
	return token.Token{
		Position: tokens[0].Position,
		Length:   token.Length(tokens[len(tokens)-1].End() - tokens[0].Position),
		Kind:     token.Type,
	}
}

// parseType tries to do speculative parsing for a type expression.
// The boolean value is true if the tokens represent a type expression.
func parseType(tokens token.List, i uint) (uint, bool) {
	for i < uint(len(tokens)) {
		t := tokens[i]

		if !allowedInType(t.Kind) {
			return i, t.Kind == token.BlockStart
		}

		i++
	}

	return i, false
}

// startsType returns true if the given `kind` can start a type expression.
func startsType(kind token.Kind) bool {
	switch kind {
	case token.Identifier, token.Mul, token.Not, token.ArrayStart:
		return true
	}

	return false
}