package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

type mockFuncRef struct {
	name string
	pkg  string
}

func (m *mockFuncRef) Name() string    { return m.name }
func (m *mockFuncRef) Package() string { return m.pkg }

func TestCall(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock("fn"))

	myfunc := &ssa.Function{
		Typ:         &types.Function{},
		FunctionRef: &mockFuncRef{name: "myfunc", pkg: "main"},
	}

	call := fn.Append(&ssa.Call{Func: myfunc, Arguments: ssa.Arguments{}})
	one := fn.Append(&ssa.Int{Int: 1})
	call2 := fn.Append(&ssa.Call{Func: myfunc, Arguments: ssa.Arguments{one}})

	assert.Equal(t, call.Type(), types.Type(types.Void))
	assert.Equal(t, call2.Type(), types.Type(types.Void))
	assert.False(t, call2.Equals(call))
}

func TestCallEquals(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock("fn"))

	sum := &ssa.Function{
		Typ: &types.Function{
			Input:  []types.Type{types.Int, types.Int},
			Output: []types.Type{types.Int},
		},
		FunctionRef: &mockFuncRef{name: "sum", pkg: "math"},
	}

	diff := &ssa.Function{
		Typ: &types.Function{
			Input:  []types.Type{types.Int, types.Int},
			Output: []types.Type{types.Int},
		},
		FunctionRef: &mockFuncRef{name: "diff", pkg: "math"},
	}

	one := fn.Append(&ssa.Int{Int: 1})
	two := fn.Append(&ssa.Int{Int: 2})
	call1 := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})
	call2 := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})
	call3 := fn.Append(&ssa.Call{Func: diff, Arguments: ssa.Arguments{one, two}})

	assert.False(t, call1.Equals(one))
	assert.True(t, call1.Equals(call2))
	assert.False(t, call1.Equals(call3))
}

func TestCallReturnType(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock("fn"))

	sum := &ssa.Function{
		Typ: &types.Function{
			Input:  []types.Type{types.Int, types.Int},
			Output: []types.Type{types.Int},
		},
	}

	one := fn.Append(&ssa.Int{Int: 1})
	two := fn.Append(&ssa.Int{Int: 2})
	call := fn.Append(&ssa.Call{Func: sum, Arguments: ssa.Arguments{one, two}})

	assert.True(t, call.Type() == types.Int)
}