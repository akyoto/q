package ssa

import (
	"slices"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Label        string
	Instructions []Value
	Predecessors []*Block
	Successors   []*Block
	Identifiers  map[string]Value
}

// NewBlock creates a new basic block.
func NewBlock(label string) *Block {
	return &Block{
		Instructions: make([]Value, 0, 8),
		Label:        label,
	}
}

// AddSuccessor adds the given block as a successor.
func (block *Block) AddSuccessor(successor *Block) {
	block.Successors = append(block.Successors, successor)
	successor.Predecessors = append(successor.Predecessors, block)
}

// Append adds a new value to the block.
func (block *Block) Append(value Value) Value {
	for _, dep := range value.Inputs() {
		dep.(HasLiveness).AddUser(value)
	}

	block.Instructions = append(block.Instructions, value)
	return value
}

// InsertAt inserts the `value` at the given `index`.
func (block *Block) InsertAt(value Value, index int) {
	for _, dep := range value.Inputs() {
		dep.(HasLiveness).AddUser(value)
	}

	block.Instructions = slices.Insert(block.Instructions, index, value)
}

// Identify adds a new identifier or changes an existing one.
func (block *Block) Identify(name string, value Value) {
	if block.Identifiers == nil {
		block.Identifiers = make(map[string]Value, 8)
	}

	block.Identifiers[name] = value
}

// RemoveNilValues removes all nil values from the block.
func (block *Block) RemoveNilValues() {
	block.Instructions = slices.DeleteFunc(block.Instructions, func(value Value) bool {
		return value == nil
	})
}

// ReplaceAll replaces all uses of `old` with `new`.
func (block *Block) ReplaceAll(old Value, new Value) {
	for _, instr := range block.Instructions {
		instr.Replace(old, new)
	}
}