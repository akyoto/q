package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	fn := ssa.Function{}
	a := fn.AppendInt(1)
	b := fn.AppendInt(2)
	c := fn.Append(ssa.Value{Type: ssa.Add, Args: []*ssa.Value{a, b}})
	fn.AddBlock()
	d := fn.AppendInt(3)
	e := fn.AppendInt(4)
	f := fn.Append(ssa.Value{Type: ssa.Add, Args: []*ssa.Value{d, e}})
	assert.Equal(t, c.String(), "1 + 2")
	assert.Equal(t, f.String(), "3 + 4")
}

func TestInvalidInstruction(t *testing.T) {
	instr := ssa.Value{}
	assert.Equal(t, instr.String(), "")
}