package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateStringOp calls a function with two strings.
func (f *Function) evaluateStringOp(op string, left ssa.Value, right ssa.Value, source ssa.Source) (ssa.Value, error) {
	concat := f.Env.Function("strings", op)
	f.Dependencies.Add(concat)
	leftStruct := left.(*ssa.Struct)
	rightStruct := right.(*ssa.Struct)

	fn := &ssa.Function{
		FunctionRef: concat,
		Typ:         concat.Type,
	}

	args := []ssa.Value{
		leftStruct.Arguments[0],
		leftStruct.Arguments[1],
		rightStruct.Arguments[0],
		rightStruct.Arguments[1],
	}

	return f.call(fn, args, source), nil
}