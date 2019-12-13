package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Assignment handles assignment instructions.
func (state *State) Assignment(tokens []token.Token) error {
	operatorPos := token.Index(tokens, token.Operator, "=")

	if operatorPos == -1 {
		return errors.New(errors.MissingAssignmentOperator)
	}

	left := tokens[:operatorPos]

	if left[operatorPos-1].Kind == token.ArrayEnd {
		return state.AssignArrayElement(tokens, operatorPos)
	}

	for _, t := range left {
		if t.Kind == token.Keyword && (t.Text() == "let" || t.Text() == "mut") {
			_, err := state.AssignVariable(tokens, false)
			return err
		}

		if t.Kind == token.Operator && t.Text() == "." {
			return state.AssignStructField(tokens, operatorPos)
		}
	}

	_, err := state.AssignVariable(tokens, false)
	return err
}
