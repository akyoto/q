package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// makeStructFromTuple creates a struct from elements in a tuple.
func (f *Function) makeStructFromTuple(tuple ssa.Value, typ types.Type, structType *types.Struct, name string, source ssa.Source) *ssa.Struct {
	fields := make([]ssa.Value, len(structType.Fields))

	for i, field := range structType.Fields {
		fieldValue := &ssa.Field{
			Tuple:  tuple,
			Index:  i,
			Source: source,
		}

		f.Block().Append(fieldValue)
		f.Block().Identify(name+"."+field.Name, fieldValue)
		fields[i] = fieldValue
	}

	return f.makeStruct(typ, fields, source)
}