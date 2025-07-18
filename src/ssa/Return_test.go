package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestReturn(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")
	ret := fn.Append(&ssa.Return{})
	one := fn.Append(&ssa.Int{Int: 1})
	ret2 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one}})
	two := fn.Append(&ssa.Int{Int: 2})
	ret3 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one, two}})
	ret4 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{two, one}})
	ret5 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one, two}})

	assert.True(t, ret.Type() == types.Void)
	assert.Equal(t, ret.String(), "return")
	assert.Equal(t, ret2.String(), "return 1")
	assert.Equal(t, ret3.String(), "return 1, 2")
	assert.Equal(t, ret4.String(), "return 2, 1")
	assert.Equal(t, ret5.String(), "return 1, 2")
	assert.False(t, ret5.Equals(one))
	assert.False(t, ret5.Equals(ret))
	assert.False(t, ret5.Equals(ret4))
	assert.True(t, ret5.Equals(ret3))
}