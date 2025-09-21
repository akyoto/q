package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// structField returns the memory for the struct field.
func (f *Function) structField(leftValue ssa.Value, field *types.Field) *ssa.Memory {
	memory, isMemory := leftValue.(*ssa.Memory)

	if isMemory {
		if memory.Scale {
			memorySize := f.Append(&ssa.Int{Int: memory.Typ.Size()})
			memory.Index = f.Append(&ssa.BinaryOp{Op: token.Mul, Left: memory.Index, Right: memorySize})
			memory.Scale = false
		}

		if field.Offset != 0 {
			offset := f.Append(&ssa.Int{Int: int(field.Offset)})
			memory.Index = f.Append(&ssa.BinaryOp{Op: token.Add, Left: memory.Index, Right: offset})
		}

		memory.Typ = field.Type
	} else {
		offset := f.Append(&ssa.Int{Int: int(field.Offset)})

		memory = &ssa.Memory{
			Address: leftValue,
			Index:   offset,
			Scale:   false,
			Typ:     field.Type,
		}
	}

	return memory
}