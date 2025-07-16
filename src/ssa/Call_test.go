package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestCall(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")
	myfunc := &ssa.Function{Name: "myfunc", Typ: &types.Function{}}
	call := fn.Append(&ssa.Call{Func: myfunc, Arguments: ssa.Arguments{}})
	one := fn.Append(&ssa.Int{Int: 1})
	call2 := fn.Append(&ssa.Call{Func: myfunc, Arguments: ssa.Arguments{one}})

	assert.True(t, call.Type() == types.Void)
	assert.Equal(t, call.String(), "myfunc()")
	assert.Equal(t, call2.String(), "myfunc(1)")
}

func TestCallEquals(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")

	sum := &ssa.Function{
		Name: "sum",
		Typ: &types.Function{
			Input:  []types.Type{types.Int, types.Int},
			Output: []types.Type{types.Int},
		},
	}

	one := fn.Append(&ssa.Int{Int: 1})
	two := fn.Append(&ssa.Int{Int: 2})
	call1 := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})
	call2 := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})

	assert.False(t, call1.Equals(one))
	assert.True(t, call1.Equals(call2))
}

func TestCallReturnType(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")

	sum := &ssa.Function{
		Name: "sum",
		Typ: &types.Function{
			Input:  []types.Type{types.Int, types.Int},
			Output: []types.Type{types.Int},
		},
	}

	one := fn.Append(&ssa.Int{Int: 1})
	two := fn.Append(&ssa.Int{Int: 2})
	call := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})

	assert.Equal(t, call.String(), "sum(1, 2)")
	assert.True(t, call.Type() == types.Int)
}