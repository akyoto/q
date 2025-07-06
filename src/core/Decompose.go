package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// Decompose creates SSA values from expressions and decomposes structs into their individual fields.
func (f *Function) Decompose(nodes []*expression.Expression, typeCheck []*ssa.Parameter) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(nodes))

	for i, node := range nodes {
		value, err := f.Evaluate(node)

		if err != nil {
			return nil, err
		}

		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			for _, field := range structure.Arguments {
				args = append(args, field)
			}

			continue
		}

		args = append(args, value)

		if typeCheck != nil && !types.Is(value.Type(), typeCheck[i].Typ) {
			_, isPointer := typeCheck[i].Typ.(*types.Pointer)

			if isPointer {
				number, isInt := value.(*ssa.Int)

				if isInt && number.Int == 0 {
					continue
				}
			}

			// Temporary hack to allow int64 -> uint32 conversion
			if types.Is(value.Type(), types.AnyInt) && types.Is(typeCheck[i].Typ, types.AnyInt) {
				continue
			}

			return nil, errors.New(&TypeMismatch{
				Encountered:   value.Type().Name(),
				Expected:      typeCheck[i].Typ.Name(),
				ParameterName: typeCheck[i].Name,
			}, f.File, value.(ssa.HasSource).Start())
		}
	}

	return args, nil
}