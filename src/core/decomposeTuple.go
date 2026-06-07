package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeTuple decomposes tuples into their individual fields with a type check.
func (f *Function) decomposeTuple(value ssa.Value, tuple *types.Tuple, typeCheck []*ssa.Parameter, source ssa.Source) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(tuple.Types))

	for i := range tuple.Types {
		if !types.Is(tuple.Types[i], typeCheck[i].Typ) {
			typeMismatch := &TypeMismatch{
				Encountered:   value.Type().Name(),
				Expected:      typeCheck[i].Typ.Name(),
				ParameterName: typeCheck[i].Name,
				IsReturn:      true,
			}

			return nil, errors.New(typeMismatch, f.File, source)
		}

		v := f.Append(&ssa.Field{
			Tuple:  value,
			Index:  i,
			Source: source,
		})

		args = append(args, v)
	}

	return args, nil
}