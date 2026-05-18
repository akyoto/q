package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateConcat concatenates two strings.
func (f *Function) evaluateConcat(left ssa.Value, right ssa.Value, source ssa.Source) (ssa.Value, error) {
	concat := f.Env.Function("strings", "concat")
	f.Dependencies.Add(concat)
	leftStruct := left.(*ssa.Struct)
	rightStruct := right.(*ssa.Struct)

	call := f.Append(&ssa.Call{
		Func: &ssa.Function{
			FunctionRef: concat,
			Typ:         concat.Type,
		},
		Arguments: []ssa.Value{leftStruct.Arguments[0], leftStruct.Arguments[1], rightStruct.Arguments[0], rightStruct.Arguments[1]},
		Source:    source,
	})

	return call, nil
}