package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// call calls a function.
func (f *Function) call(fn *ssa.Function, args []ssa.Value, source ssa.Source) ssa.Value {
	call := f.Append(&ssa.Call{
		Func:      fn,
		Arguments: args,
		Source:    source,
	})

	f.Dependencies.Add(fn.FunctionRef.(*Function))
	typ := call.Type()
	structure, isStructType := types.Unwrap(typ).(*types.Struct)

	if isStructType {
		var fields []ssa.Value

		for _, field := range structure.Fields {
			ssaField := f.Append(&ssa.Field{
				Tuple: call,
				Index: int(field.Index),
			})

			fields = append(fields, ssaField)
		}

		return f.makeStruct(typ, fields, source)
	}

	return call
}