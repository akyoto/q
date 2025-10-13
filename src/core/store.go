package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// store stores a value at the given memory address.
func (f *Function) store(memory *ssa.Memory, value ssa.Value) error {
	if !types.Is(value.Type(), memory.Typ) {
		return errors.New(&TypeMismatch{Encountered: value.Type().Name(), Expected: memory.Typ.Name()}, f.File, value.(errors.Source))
	}

	structure, isStructType := types.Unwrap(value.Type()).(*types.Struct)

	if !isStructType {
		f.Append(&ssa.Store{
			Memory: memory,
			Value:  value,
		})

		return nil
	}

	composite, isStruct := value.(*ssa.Struct)

	if isStruct {
		for i, field := range structure.Fields {
			fieldValue := composite.Arguments[i]

			f.Append(&ssa.Store{
				Memory: f.structField(memory, field),
				Value:  fieldValue,
			})
		}

		return nil
	}

	fieldValues := make([]ssa.Value, len(structure.Fields))

	for i := range structure.Fields {
		fieldValue := &ssa.FromTuple{
			Tuple:  value,
			Index:  i,
			Source: token.NewSource(value.(errors.Source).Start(), value.(errors.Source).End()),
		}

		f.Block().Append(fieldValue)
		fieldValues[i] = fieldValue
	}

	for i, field := range structure.Fields {
		f.Append(&ssa.Store{
			Memory: f.structField(memory, field),
			Value:  fieldValues[i],
		})
	}

	return nil
}