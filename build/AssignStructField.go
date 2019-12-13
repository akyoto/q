package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// AssignStructField assigns a value to a struct field.
func (state *State) AssignStructField(tokens []token.Token, operatorPos token.Position) error {
	left := tokens[:operatorPos]
	variableName := left[0].Text()
	fieldName := left[2].Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		return errors.New(state.UnknownVariableError(variableName))
	}

	field := variable.Type.FieldByName(fieldName)

	if field == nil {
		return errors.New(UnknownFieldError(fieldName, variable.Type))
	}

	right := tokens[operatorPos+1:]

	if len(right) == 1 && right[0].Kind == token.Number {
		number, err := state.ParseInt(right[0].Text())

		if err != nil {
			return errors.New(err)
		}

		state.assembler.StoreNumber(variable.Register(), byte(field.Offset), byte(field.Type.Size), uint64(number))
	} else {
		return errors.New(errors.NotImplemented)
	}

	return nil
}
