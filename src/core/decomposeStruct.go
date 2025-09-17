package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeStruct packs struct fields into registers and adds them to `args`.
func (f *Function) decomposeStruct(args []ssa.Value, structure *ssa.Struct) []ssa.Value {
	if structure.Typ.Size() > 16 {
		for _, field := range structure.Arguments {
			args = append(args, field)
		}

		return args
	}

	// Packed integer: Use the first argument,
	// then bitwise OR with the shifted field values.
	cursor := structure.Arguments[0]
	typ := types.Unwrap(structure.Typ).(*types.Struct)
	size := typ.Fields[0].Type.Size()

	for i, field := range structure.Arguments[1:] {
		fieldSize := typ.Fields[i+1].Type.Size()

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
	return args
}