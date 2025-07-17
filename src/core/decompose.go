package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decompose creates SSA values from expressions and decomposes structs into their individual fields.
func (f *Function) decompose(nodes []*expression.Expression, typeCheck []*ssa.Parameter) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(nodes))

	for i, node := range nodes {
		value, err := f.eval(node)

		if err != nil {
			return nil, err
		}

		if typeCheck != nil && !types.Is(value.Type(), typeCheck[i].Typ) {
			return nil, errors.New(&TypeMismatch{
				Encountered:   value.Type().Name(),
				Expected:      typeCheck[i].Typ.Name(),
				ParameterName: typeCheck[i].Name,
			}, f.File, value.(ssa.HasSource).Start())
		}

		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			for _, field := range structure.Arguments {
				args = append(args, field)
			}

			continue
		}

		args = append(args, value)
	}

	return args, nil
}