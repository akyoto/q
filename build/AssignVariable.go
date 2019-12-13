package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// AssignVariable handles assignment instructions and also returns the referenced variable.
func (state *State) AssignVariable(tokens []token.Token, isNewVariable bool) (*Variable, error) {
	cursor := 0
	mutable := false
	left := tokens[cursor]

	if left.Kind == token.Keyword {
		switch left.Text() {
		case "let":
			isNewVariable = true

		case "mut":
			mutable = true
			isNewVariable = true

		default:
			return nil, errors.New(errors.InvalidExpression)
		}

		cursor++
		state.tokenCursor++
		left = tokens[cursor]
	}

	if left.Kind != token.Identifier {
		return nil, errors.New(errors.ExpectedVariable)
	}

	if tokens[cursor+1].Kind != token.Operator {
		return nil, errors.New(errors.MissingAssignmentOperator)
	}

	assignPos := state.tokenCursor
	variableName := left.Text()
	variable := state.scopes.Get(variableName)

	if isNewVariable {
		if variable != nil {
			return variable, errors.New(&errors.VariableAlreadyExists{Name: variable.Name})
		}

		register := state.registers.General.FindFree()

		if register == nil {
			return nil, errors.ExceededMaxVariables
		}

		variable = &Variable{
			Name:           variableName,
			Position:       assignPos,
			LastAssign:     assignPos,
			LastAssignUsed: false,
			Mutable:        mutable,
			AliveUntil:     state.identifierLifeTime[variableName],
		}

		variable.ForceSetRegister(register)
		defer state.scopes.Add(variable)
	} else {
		if variable == nil {
			return nil, errors.New(state.UnknownVariableError(variableName))
		}

		if !variable.Mutable {
			return variable, errors.New(&errors.ImmutableVariable{Name: variable.Name})
		}
	}

	// Skip operator
	cursor++
	state.tokenCursor++

	// Expression
	cursor++
	state.tokenCursor++
	value := tokens[cursor:]

	if len(value) == 0 {
		return variable, errors.MissingAssignmentExpression
	}

	// A question token indicates an unknown value.
	if len(value) == 1 && value[0].Kind == token.Question {
		variable.LastAssignUsed = true
		state.tokenCursor += len(value)
		return variable, nil
	}

	// Move result of expression to register
	typ, err := state.TokensToRegister(value, variable.Register())

	if err != nil {
		return variable, err
	}

	if isNewVariable {
		variable.Type = typ
	} else if typ != variable.Type {
		return variable, errors.New(&errors.InvalidType{Type: typ.String(), Expected: variable.Type.String()})
	}

	// Check for ineffective assignments
	if !isNewVariable {
		if !variable.LastAssignUsed {
			state.tokenCursor = variable.LastAssign
			return variable, errors.New(&errors.IneffectiveAssignment{Name: variable.Name})
		}

		variable.LastAssign = assignPos
	}

	// Currently we can't prove that the value hasn't been used inside a loop.
	if !state.InLoop() {
		variable.LastAssignUsed = false
	}

	state.tokenCursor += len(value)
	return variable, nil
}
