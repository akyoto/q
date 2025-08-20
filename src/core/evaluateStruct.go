package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateStruct converts a struct expression to an SSA value.
func (f *Function) evaluateStruct(expr *expression.Expression) (ssa.Value, error) {
	typ := ParseType([]token.Token{expr.Children[0].Token}, f.File.Bytes, f.Env).(*types.Struct)

	structure := &ssa.Struct{
		Typ:       typ,
		Arguments: make(ssa.Arguments, len(typ.Fields)),
	}

	for _, definition := range expr.Children[1:] {
		if len(definition.Children) != 2 {
			return nil, errors.New(InvalidFieldInit, f.File, definition.Source().StartPos)
		}

		left := definition.Children[0]
		fieldName := left.String(f.File.Bytes)
		field := typ.FieldByName(fieldName)
		right := definition.Children[1]
		rightValue, err := f.evaluate(right)

		if err != nil {
			return nil, err
		}

		structure.Arguments[field.Index] = rightValue
	}

	return structure, nil
}