package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// composeStruct unpacks registers into struct fields.
func (f *Function) composeStruct(structure *ssa.Struct, structType *types.Struct, input *ssa.Parameter, i int, offset int) int {
	if structType.Size() > 16 {
		for _, field := range structType.Fields {
			param := &ssa.Parameter{
				Index:     uint8(offset + i),
				Name:      input.Name + "." + field.Name,
				Typ:       field.Type,
				Tokens:    input.Tokens,
				Structure: structure,
				Source:    input.Source,
			}

			f.Append(param)
			f.Block().Identify(param.Name, param)
			structure.Arguments = append(structure.Arguments, param)
			offset++
		}

		return offset - 1
	}

	var (
		param *ssa.Parameter
		size  = 8
	)

	for _, field := range structType.Fields {
		fieldSize := structType.Fields[i].Type.Size()

		if size+fieldSize > 8 {
			param = &ssa.Parameter{
				Index:     uint8(offset + i),
				Typ:       field.Type,
				Tokens:    input.Tokens,
				Structure: structure,
				Source:    input.Source,
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
				sizeValue := f.Append(&ssa.Int{Int: size * 8})
				shifted = f.Append(&ssa.BinaryOp{Op: token.Shr, Left: param, Right: sizeValue})
			} else {
				shifted = param
			}

			mask := f.Append(&ssa.Int{Int: (1 << (fieldSize * 8)) - 1})
			fieldValue = f.Append(&ssa.BinaryOp{Op: token.And, Left: shifted, Right: mask, Source: param.Source, Structure: structure})
		}

		f.Block().Identify(input.Name+"."+field.Name, fieldValue)
		structure.Arguments = append(structure.Arguments, fieldValue)
		size += fieldSize
	}

	return offset - 1
}