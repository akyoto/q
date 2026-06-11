package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateStruct converts a struct expression to an SSA value.
func (f *Function) evaluateStruct(expr *expression.Expression) (ssa.Value, error) {
	if expr.Children[0].Token.Kind == token.Call && expr.Children[0].Children[0].Token.Kind == token.New {
		return f.evaluateNewStruct(expr)
	}

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
		Source:    expr.Source(),
	}

	for _, definition := range expr.Children[1:] {
		if isTrailing(definition, expr.Children) {
			continue
		}

		field, rightValue, err := f.extractField(structType, definition)

		if err != nil {
			return nil, err
		}

		structure.Arguments[field.Index] = rightValue
	}

	return structure, nil
}