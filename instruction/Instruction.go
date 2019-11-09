package instruction

import "github.com/akyoto/q/token"

// Instruction encapsulates a single instruction inside a function.
// Instructions can be variable assignments, function calls or keywords.
type Instruction struct {
	Kind     Kind
	Tokens   []token.Token
	Position token.Position
}
