package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// composeStruct unpacks registers into struct fields.
func (f *Function) composeStruct(structType *types.Struct, input *ssa.Parameter, i int, offset int) (*ssa.Struct, int) {
	fields := make([]ssa.Value, 0, len(structType.Fields))

	if structType.Size() > 16 {
		for _, field := range structType.Fields {
			param := &ssa.Parameter{
				Index:  uint8(offset + i),
				Name:   input.Name + "." + field.Name,
				Typ:    field.Type,
				Tokens: input.Tokens,
				Source: input.Source,
			}

			f.Append(param)
			f.Block().Identify(param.Name, param)
			fields = append(fields, param)
			offset++
		}

		structure := f.makeStruct(structType, input.Source, fields)
		return structure, offset - 1
	}

	var (
		param *ssa.Parameter
		size  = 8
	)

	for _, field := range structType.Fields {
		fieldSize := structType.Fields[i].Type.Size()

		if size+fieldSize > 8 {
			param = &ssa.Parameter{
				Index:  uint8(offset + i),
				Typ:    field.Type,
				Tokens: input.Tokens,
				Source: input.Source,
			}

			f.Append(param)
			offset++
			size = 0
		}

		var fieldValue ssa.Value

		if size == 0 && fieldSize == 8 {
			fieldValue = param
		} else {
			var shifted ssa.Value

			if size > 0 {
				param.Typ = types.UInt
				sizeValue := f.Append(&ssa.Int{Int: size * 8})
				shifted = f.Append(&ssa.BinaryOp{Op: token.Shr, Left: param, Right: sizeValue})
			} else {
				shifted = param
			}

			mask := f.Append(&ssa.Int{Int: (1 << (fieldSize * 8)) - 1})
			fieldValue = f.Append(&ssa.BinaryOp{Op: token.And, Left: shifted, Right: mask, Source: param.Source})
		}

		f.Block().Identify(input.Name+"."+field.Name, fieldValue)
		fields = append(fields, fieldValue)
		size += fieldSize
	}

	structure := f.makeStruct(structType, input.Source, fields)
	return structure, offset - 1
}