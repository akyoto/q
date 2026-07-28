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

	f.Calls.Add(fn.FunctionRef.(*Function))
	typ := call.Type()
	structure, isStructType := types.Unwrap(typ).(*types.Struct)

	if isStructType {
		fields := make([]ssa.Value, len(structure.Fields))

		for i := range fields {
			fields[i] = f.Append(&ssa.Field{
				Tuple: call,
				Index: i,
			})
		}

		return f.makeStruct(typ, fields, source)
	}

	return call
}