package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Assignment handles assignment instructions.
func (state *State) Assignment(tokens []token.Token) error {
	_, err := state.AssignVariable(tokens)
	return err
}

// AssignVariable handles assignment instructions and also returns the referenced variable.
func (state *State) AssignVariable(tokens []token.Token) (*Variable, error) {
	cursor := 0
	mutable := false
	left := tokens[cursor]

	if left.Kind == token.Keyword && left.Text() == "mut" {
		mutable = true
		cursor++
		state.tokenCursor++
		left = tokens[cursor]
	}

	if left.Kind != token.Identifier {
		return nil, errors.ExpectedVariable
	}

	variableName := left.Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		variable = &Variable{
			Name:     variableName,
			Position: state.tokenCursor,
			Mutable:  mutable,
		}

		state.scopes.Add(variable)
	} else if !variable.Mutable {
		return variable, &errors.ImmutableVariable{VariableName: variable.Name}
	}

	// Operator
	cursor++
	state.tokenCursor++
	operator := tokens[cursor]

	if operator.Kind != token.Operator {
		return variable, errors.MissingAssignmentOperator
	}

	// Expression
	cursor++
	state.tokenCursor++
	value := tokens[cursor:]

	if len(value) == 0 {
		return variable, errors.MissingAssignmentExpression
	}

	reg, err := state.EvaluateTokens(value)

	if err != nil {
		return variable, err
	}

	variable.ForceSetRegister(reg)
	state.tokenCursor += len(value)
	return variable, nil
}
