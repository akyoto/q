package ssa

import (
	"slices"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Identifiers  map[string]Value
	Label        string
	Instructions []Value
	Predecessors []*Block
	Successors   []*Block
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
func (block *Block) Append(value Value) {
	block.Instructions = append(block.Instructions, value)
}

// InsertAt inserts the `value` at the given `index`.
func (block *Block) InsertAt(value Value, index int) {
	block.Instructions = slices.Insert(block.Instructions, index, value)
}

// Identify adds a new identifier or changes an existing one.
func (block *Block) Identify(name string, value Value) {
	if block.Identifiers == nil {
		block.Identifiers = make(map[string]Value, 8)
	}

	block.Identifiers[name] = value
}

// LookupIdentifier searches for the possible definitions the identifier can have.
func (block *Block) LookupIdentifier(name string, traversed map[*Block]bool) (value Value, appended bool) {
	if traversed[block] {
		return nil, false
	}

	traversed[block] = true
	value, exists := block.Identifiers[name]

	if exists {
		return value, true
	}

	switch len(block.Predecessors) {
	case 0:
		return nil, false
	case 1:
		return block.Predecessors[0].LookupIdentifier(name, traversed)
	default:
		var values []Value

		for _, pre := range block.Predecessors {
			value, exists := pre.LookupIdentifier(name, traversed)

			if exists {
				values = append(values, value)
			}
		}

		if len(values) == 0 {
			return nil, false
		}

		if len(values) == 1 {
			return values[0], true
		}

		return &Phi{Arguments: values}, false
	}
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