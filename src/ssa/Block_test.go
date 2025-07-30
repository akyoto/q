package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/go/assert"
)

func TestBlockCanReachPredecessor(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	a.AddSuccessor(b)

	assert.True(t, a.CanReachPredecessor(a))
	assert.True(t, b.CanReachPredecessor(a))
	assert.True(t, b.CanReachPredecessor(b))
	assert.False(t, a.CanReachPredecessor(b))
}

func TestBlockCanReachPredecessorBranch(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	c := ssa.NewBlock("c")
	d := ssa.NewBlock("d")
	a.AddSuccessor(b)
	a.AddSuccessor(c)
	b.AddSuccessor(d)
	c.AddSuccessor(d)

	assert.True(t, a.CanReachPredecessor(a))
	assert.True(t, b.CanReachPredecessor(a))
	assert.True(t, c.CanReachPredecessor(a))
	assert.True(t, d.CanReachPredecessor(a))
	assert.True(t, d.CanReachPredecessor(b))
	assert.True(t, d.CanReachPredecessor(c))
	assert.False(t, b.CanReachPredecessor(c))
	assert.False(t, c.CanReachPredecessor(b))
	assert.False(t, b.CanReachPredecessor(d))
	assert.False(t, c.CanReachPredecessor(d))
}

func TestBlockCanReachPredecessorLoop(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	c := ssa.NewBlock("c")
	a.AddSuccessor(b)
	b.AddSuccessor(b)
	b.AddSuccessor(c)

	assert.True(t, a.CanReachPredecessor(a))
	assert.False(t, a.CanReachPredecessor(b))
	assert.False(t, a.CanReachPredecessor(c))
	assert.True(t, b.CanReachPredecessor(a))
	assert.True(t, b.CanReachPredecessor(b))
	assert.False(t, b.CanReachPredecessor(c))
	assert.True(t, c.CanReachPredecessor(a))
	assert.True(t, c.CanReachPredecessor(b))
	assert.True(t, c.CanReachPredecessor(c))
}