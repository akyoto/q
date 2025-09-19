package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decompose creates SSA values from expressions and decomposes structs into their individual fields.
func (f *Function) decompose(nodes []*expression.Expression, typeCheck []*ssa.Parameter, isReturn bool) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(nodes))

	for i, node := range nodes {
		value, err := f.evaluateRight(node)

		if err != nil {
			return nil, err
		}

		_, isSyscall := value.(*ssa.Syscall)

		if typeCheck != nil && !isSyscall {
			valueType := value.Type()
			expectedType := typeCheck[i].Typ
			_, valueIsResource := valueType.(*types.Resource)
			expectedResource, expectedIsResource := expectedType.(*types.Resource)

			if valueIsResource && expectedIsResource {
				f.Block().Unidentify(value)
			}

			if isReturn && expectedIsResource && types.Is(valueType, expectedResource.Of) {
				// pass type check.
			} else if !types.Is(valueType, expectedType) {
				typeMismatch := &TypeMismatch{
					Encountered:   value.Type().Name(),
					Expected:      typeCheck[i].Typ.Name(),
					ParameterName: typeCheck[i].Name,
					IsReturn:      isReturn,
				}

				return nil, errors.New(typeMismatch, f.File, node.Source().StartPos)
			}
		}

		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			args = f.decomposeStruct(args, structure)
			continue
		}

		args = append(args, value)
	}

	return args, nil
}