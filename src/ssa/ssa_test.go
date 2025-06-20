package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
)

func TestBlock(t *testing.T) {
	block := &ssa.Block{}
	block.Append(ssa.Instruction{Type: ssa.Int, Int: 1})
	block.Append(ssa.Instruction{Type: ssa.Int, Int: 2})
	block.Append(ssa.Instruction{Type: ssa.Add, Parameters: []ssa.Index{0, 1}})
}