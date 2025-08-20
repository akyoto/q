package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// decompose creates SSA values from expressions and decomposes structs into their individual fields.
func (f *Function) decompose(nodes []*expression.Expression, typeCheck []*ssa.Parameter, isReturn bool) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(nodes))

	for i, node := range nodes {
		value, err := f.evaluate(node)

		if err != nil {
			return nil, err
		}

		if typeCheck != nil && !types.Is(value.Type(), typeCheck[i].Typ) {
			return nil, errors.New(&TypeMismatch{
				Encountered:   value.Type().Name(),
				Expected:      typeCheck[i].Typ.Name(),
				ParameterName: typeCheck[i].Name,
				IsReturn:      isReturn,
			}, f.File, node.Source().StartPos)
		}

		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			switch {
			case structure.Typ.Size() <= 16:
				// Packed integer: Use the first argument,
				// then bitwise OR with the shifted field values.
				cursor := structure.Arguments[0]
				size := structure.Typ.Fields[0].Type.Size()

				for i, field := range structure.Arguments[1:] {
					fieldSize := structure.Typ.Fields[i+1].Type.Size()

					if size+fieldSize > 8 {
						// The field doesn't fit into the register anymore.
						// We need to use this field as the starting value
						// for the next argument.
						args = append(args, cursor)
						cursor = field
						size = fieldSize
						continue
					}

					sizeValue := f.Append(&ssa.Int{Int: size * 8})
					shifted := f.Append(&ssa.BinaryOp{Op: token.Shl, Left: field, Right: sizeValue})
					cursor = f.Append(&ssa.BinaryOp{Op: token.Or, Left: cursor, Right: shifted})
					size += fieldSize
				}

				args = append(args, cursor)
			default:
				for _, field := range structure.Arguments {
					args = append(args, field)
				}
			}

			continue
		}

		args = append(args, value)
	}

	return args, nil
}