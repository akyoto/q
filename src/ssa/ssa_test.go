package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	fn := ssa.IR{}
	a := fn.AppendInt(1)
	b := fn.AppendInt(2)
	c := fn.Append(&ssa.BinaryOperation{Op: token.Add, Left: a, Right: b})
	fn.AddBlock()
	d := fn.AppendInt(3)
	e := fn.AppendInt(4)
	f := fn.Append(&ssa.BinaryOperation{Op: token.Add, Left: d, Right: e})
	assert.Equal(t, c.String(), "1 + 2")
	assert.Equal(t, f.String(), "3 + 4")
}