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
	typ, err := f.Env.TypeFromTokens(token.List{expr.Children[0].Token}, f.File)

	if err != nil {
		return nil, err
	}

	structType, isStructType := typ.(*types.Struct)

	if !isStructType {
		panic("not a struct")
	}

	structure := &ssa.Struct{
		Typ:       structType,
		Arguments: make(ssa.Arguments, len(structType.Fields)),
	}

	for _, definition := range expr.Children[1:] {
		if len(definition.Children) != 2 {
			return nil, errors.New(InvalidFieldInit, f.File, definition.Source().StartPos)
		}

		left := definition.Children[0]

		if left.Token.Kind != token.Identifier {
			if left.Token.Kind == token.FieldAssign {
				return nil, errors.New(MissingCommaBetweenFields, f.File, left.Source().StartPos)
			}

			return nil, errors.New(InvalidFieldInit, f.File, left.Source().StartPos)
		}

		fieldName := left.String(f.File.Bytes)
		field := structType.FieldByName(fieldName)

		if field == nil {
			return nil, errors.New(&UnknownStructField{StructName: typ.Name(), FieldName: fieldName}, f.File, left.Source().StartPos)
		}

		right := definition.Children[1]
		rightValue, err := f.evaluate(right)

		if err != nil {
			return nil, err
		}

		structure.Arguments[field.Index] = rightValue
	}

	return structure, nil
}