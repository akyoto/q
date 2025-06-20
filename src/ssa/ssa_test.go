package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/go/assert"
)

func TestBlock(t *testing.T) {
	f := ssa.Function{}
	block := f.AddBlock()
	a := block.Append(ssa.Instruction{Type: ssa.Int, Int: 1})
	b := block.Append(ssa.Instruction{Type: ssa.Int, Int: 2})
	c := block.Append(ssa.Instruction{Type: ssa.Add, Args: []*ssa.Instruction{a, b}})
	assert.Equal(t, c.String(), "1 + 2")
}

func TestInvalidInstruction(t *testing.T) {
	instr := ssa.Instruction{}
	assert.Equal(t, instr.String(), "")
}