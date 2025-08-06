package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestReturn(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock("fn"))
	ret := fn.Append(&ssa.Return{})
	one := fn.Append(&ssa.Int{Int: 1})
	ret2 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one}})
	two := fn.Append(&ssa.Int{Int: 2})
	ret3 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one, two}})
	ret4 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{two, one}})
	ret5 := fn.Append(&ssa.Return{Arguments: ssa.Arguments{one, two}})

	assert.True(t, ret.Type() == types.Void)
	assert.False(t, ret2.Equals(ret))
	assert.False(t, ret5.Equals(one))
	assert.False(t, ret5.Equals(ret))
	assert.False(t, ret5.Equals(ret4))
	assert.True(t, ret5.Equals(ret3))
}