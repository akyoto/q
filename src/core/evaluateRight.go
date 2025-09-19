package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateRight creates a load from memory if the value
// was a memory address, otherwise returns the raw value.
func (f *Function) evaluateRight(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluate(expr)

	if err != nil {
		return nil, err
	}

	memory, isMemory := value.(*ssa.Memory)

	if !isMemory {
		return value, nil
	}

	switch typ := memory.Typ.(type) {
	case *types.Struct:
		fields := make([]ssa.Value, 0, len(typ.Fields))

		for _, field := range typ.Fields {
			fieldMemory := f.Append(&ssa.Memory{
				Address: memory.Address,
				Index:   f.Append(&ssa.Int{Int: int(field.Offset)}),
				Scale:   false,
				Typ:     field.Type,
			})

			fieldValue := f.Append(&ssa.Load{
				Memory: fieldMemory,
				Source: expr.Source(),
			})

			fields = append(fields, fieldValue)
		}

		value := &ssa.Struct{
			Typ:       typ,
			Arguments: fields,
			Source:    expr.Source(),
		}

		return value, nil

	default:
		load := f.Append(&ssa.Load{
			Memory: memory,
			Source: memory.Source,
		})

		return load, nil
	}
}