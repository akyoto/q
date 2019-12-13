package build

import (
	"fmt"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/log"
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
		return state.AssignArrayElement(tokens)
	}

	for _, t := range left {
		if t.Kind == token.Keyword && (t.Text() == "let" || t.Text() == "mut") {
			_, err := state.AssignVariable(tokens, false)
			return err
		}

		if t.Kind == token.Operator && t.Text() == "." {
			return state.AssignStructField(tokens)
		}
	}

	_, err := state.AssignVariable(tokens, false)
	return err
}

// AssignStructField assigns a value to a struct field.
func (state *State) AssignStructField(tokens []token.Token) error {
	log.Info.Println("AssignStructField", tokens)
	return nil
}

// AssignArrayElement assigns a value to an array element.
func (state *State) AssignArrayElement(tokens []token.Token) error {
	log.Info.Println("AssignArrayElement", tokens)
	operatorPos := token.Index(tokens, token.Operator, "=")
	left := tokens[:operatorPos]
	suffix := left[1:]

	if suffix[0].Kind == token.ArrayStart && suffix[len(suffix)-1].Kind == token.ArrayEnd {
		indexTokens := suffix[1 : len(suffix)-1]

		if indexTokens[0].Kind != token.Number {
			return errors.New(errors.NotImplemented)
		}

		fmt.Println(indexTokens)
		return errors.New(errors.NotImplemented)
	}

	return nil
}

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
