package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/go/assert"
)

func TestBlockCanReachSelf(t *testing.T) {
	a := ssa.NewBlock("a")
	assert.True(t, a.CanReachPredecessor(a))
}

func TestBlockCanReachPredecessor(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	a.AddSuccessor(b)

	assert.False(t, a.CanReachPredecessor(b))
	assert.True(t, b.CanReachPredecessor(a))
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

	assert.False(t, a.CanReachPredecessor(b))
	assert.False(t, a.CanReachPredecessor(c))
	assert.False(t, a.CanReachPredecessor(d))

	assert.True(t, b.CanReachPredecessor(a))
	assert.False(t, b.CanReachPredecessor(c))
	assert.False(t, b.CanReachPredecessor(d))

	assert.True(t, c.CanReachPredecessor(a))
	assert.False(t, c.CanReachPredecessor(b))
	assert.False(t, c.CanReachPredecessor(d))

	assert.True(t, d.CanReachPredecessor(a))
	assert.True(t, d.CanReachPredecessor(b))
	assert.True(t, d.CanReachPredecessor(c))
}

func TestBlockCanReachPredecessorLoop(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	c := ssa.NewBlock("c")
	a.AddSuccessor(b)
	b.AddSuccessor(b)
	b.AddSuccessor(c)

	assert.False(t, a.CanReachPredecessor(b))
	assert.False(t, a.CanReachPredecessor(c))

	assert.True(t, b.CanReachPredecessor(a))
	assert.False(t, b.CanReachPredecessor(c))

	assert.True(t, c.CanReachPredecessor(a))
	assert.True(t, c.CanReachPredecessor(b))
}

func TestBlockCanReachPredecessorLoopExtended(t *testing.T) {
	a := ssa.NewBlock("a")
	b := ssa.NewBlock("b")
	c := ssa.NewBlock("c")
	d := ssa.NewBlock("d")
	a.AddSuccessor(b)
	b.AddSuccessor(c)
	b.AddSuccessor(d)
	c.AddSuccessor(b)

	assert.False(t, a.CanReachPredecessor(b))
	assert.False(t, a.CanReachPredecessor(c))
	assert.False(t, a.CanReachPredecessor(d))

	assert.True(t, b.CanReachPredecessor(a))
	assert.True(t, b.CanReachPredecessor(c))
	assert.False(t, b.CanReachPredecessor(d))

	assert.True(t, c.CanReachPredecessor(a))
	assert.True(t, c.CanReachPredecessor(b))
	assert.False(t, c.CanReachPredecessor(d))

	assert.True(t, d.CanReachPredecessor(a))
	assert.True(t, d.CanReachPredecessor(b))
	assert.True(t, d.CanReachPredecessor(c))
}