package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestBinaryOp(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")
	a := fn.Append(&ssa.Int{Int: 1})
	b := fn.Append(&ssa.Int{Int: 2})
	c := fn.Append(&ssa.BinaryOp{Op: token.Add, Left: a, Right: b})
	fn.AddBlock("")
	d := fn.Append(&ssa.Int{Int: 3})
	e := fn.Append(&ssa.Int{Int: 4})
	f := fn.Append(&ssa.BinaryOp{Op: token.Add, Left: d, Right: e})

	assert.Equal(t, c.String(), "1 + 2")
	assert.Equal(t, f.String(), "3 + 4")
	assert.True(t, c.Type() == types.AnyInt)
}

func TestBinaryOpEquals(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock("")
	one := fn.Append(&ssa.Int{Int: 1})
	two := fn.Append(&ssa.Int{Int: 2})
	binOp := fn.Append(&ssa.BinaryOp{Op: token.Add, Left: one, Right: two})
	oneDup := fn.Append(&ssa.Int{Int: 1})
	twoDup := fn.Append(&ssa.Int{Int: 2})
	binOpDup := fn.Append(&ssa.BinaryOp{Op: token.Add, Left: oneDup, Right: twoDup})
	binOpDiff := fn.Append(&ssa.BinaryOp{Op: token.Add, Left: oneDup, Right: oneDup})

	assert.False(t, one.Equals(two))
	assert.False(t, one.Equals(binOp))
	assert.True(t, one.Equals(oneDup))
	assert.False(t, two.Equals(one))
	assert.False(t, two.Equals(binOp))
	assert.True(t, two.Equals(twoDup))
	assert.False(t, binOp.Equals(binOpDiff))
	assert.True(t, binOp.Equals(binOpDup))
}