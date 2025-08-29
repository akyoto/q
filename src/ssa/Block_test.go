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

func TestBlockFindIdentifier(t *testing.T) {
	branch := ssa.NewBlock("branch")
	thenBlock := ssa.NewBlock("if.then")
	elseBlock := ssa.NewBlock("if.else")
	mergeBlock := ssa.NewBlock("merge")

	branchValue := ssa.Value(&ssa.Int{Int: 1})
	branch.Append(branchValue)
	branch.Identify("branch", branchValue)
	branch.Identify("x", branchValue)
	branch.AddSuccessor(thenBlock)
	branch.AddSuccessor(elseBlock)

	thenValue := ssa.Value(&ssa.Int{Int: 2})
	thenBlock.Append(thenValue)
	thenBlock.Identify("then", thenValue)
	thenBlock.Identify("x", thenValue)
	thenBlock.AddSuccessor(mergeBlock)

	elseValue := ssa.Value(&ssa.Int{Int: 3})
	elseBlock.Append(elseValue)
	elseBlock.Identify("else", elseValue)
	elseBlock.Identify("x", elseValue)
	elseBlock.AddSuccessor(mergeBlock)

	mergeValue := ssa.Value(&ssa.Int{Int: 4})
	mergeBlock.Append(mergeValue)
	mergeBlock.Identify("merge", mergeValue)

	// Branch
	value, exists := branch.FindIdentifier("branch")
	assert.True(t, exists)
	assert.Equal(t, value, branchValue)

	value, exists = thenBlock.FindIdentifier("branch")
	assert.True(t, exists)
	assert.Equal(t, value, branchValue)

	value, exists = elseBlock.FindIdentifier("branch")
	assert.True(t, exists)
	assert.Equal(t, value, branchValue)

	value, exists = mergeBlock.FindIdentifier("branch")
	assert.True(t, exists)
	assert.Equal(t, value, branchValue)

	// Then
	_, exists = branch.FindIdentifier("then")
	assert.False(t, exists)

	value, exists = thenBlock.FindIdentifier("then")
	assert.True(t, exists)
	assert.Equal(t, value, thenValue)

	_, exists = elseBlock.FindIdentifier("then")
	assert.False(t, exists)

	partial, exists := mergeBlock.FindIdentifier("then")
	assert.True(t, exists)
	phi, isPhi := partial.(*ssa.Phi)
	assert.True(t, isPhi)
	assert.True(t, phi.IsPartiallyUndefined())
	assert.Equal(t, phi.FirstDefined(), value)

	// Else
	_, exists = branch.FindIdentifier("else")
	assert.False(t, exists)

	_, exists = thenBlock.FindIdentifier("else")
	assert.False(t, exists)

	value, exists = elseBlock.FindIdentifier("else")
	assert.True(t, exists)
	assert.Equal(t, value, elseValue)

	partial, exists = mergeBlock.FindIdentifier("else")
	assert.True(t, exists)
	phi, isPhi = partial.(*ssa.Phi)
	assert.True(t, isPhi)
	assert.True(t, phi.IsPartiallyUndefined())
	assert.Equal(t, phi.FirstDefined(), value)

	// Merge
	_, exists = branch.FindIdentifier("merge")
	assert.False(t, exists)

	_, exists = thenBlock.FindIdentifier("merge")
	assert.False(t, exists)

	_, exists = elseBlock.FindIdentifier("merge")
	assert.False(t, exists)

	value, exists = mergeBlock.FindIdentifier("merge")
	assert.True(t, exists)
	assert.Equal(t, value, mergeValue)

	// Phi
	value, exists = mergeBlock.FindIdentifier("x")
	assert.True(t, exists)
	phi, isPhi = value.(*ssa.Phi)
	assert.True(t, isPhi)
	assert.Equal(t, phi.Arguments[0], thenValue)
	assert.Equal(t, phi.Arguments[1], elseValue)
}

func TestBlockIndex(t *testing.T) {
	a := ssa.NewBlock("a")
	one := ssa.Value(&ssa.Int{Int: 1})
	two := ssa.Value(&ssa.Int{Int: 2})
	three := ssa.Value(&ssa.Int{Int: 3})
	other := ssa.Value(&ssa.Int{Int: 42})

	a.Append(one)
	a.Append(two)
	a.Append(three)

	assert.Equal(t, a.Index(one), 0)
	assert.Equal(t, a.Index(two), 1)
	assert.Equal(t, a.Index(three), 2)
	assert.Equal(t, a.Index(other), -1)
}