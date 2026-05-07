package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// extractField extracts the struct field and the value it is supposed to be initialized with.
func (f *Function) extractField(structType *types.Struct, definition *expression.Expression) (*types.Field, ssa.Value, error) {
	if len(definition.Children) != 2 {
		return nil, nil, errors.New(InvalidFieldInit, f.File, definition.Source())
	}

	left := definition.Children[0]

	if left.Token.Kind != token.Identifier {
		if left.Token.Kind == token.FieldAssign {
			return nil, nil, errors.New(MissingCommaBetweenFields, f.File, left.Source())
		}

		return nil, nil, errors.New(InvalidFieldInit, f.File, left.Source())
	}

	fieldName := left.String(f.File.Bytes)
	field := structType.FieldByName(fieldName)

	if field == nil {
		return nil, nil, errors.New(&UnknownStructField{StructName: structType.Name(), FieldName: fieldName}, f.File, left.Source())
	}

	right := definition.Children[1]
	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return nil, nil, err
	}

	if !types.Is(rightValue.Type(), field.Type) {
		return nil, nil, errors.New(&TypeMismatch{Encountered: rightValue.Type().Name(), Expected: field.Type.Name()}, f.File, right.Source())
	}

	return field, rightValue, nil
}