package ssa

import (
	"slices"

	"git.urbach.dev/cli/q/src/set"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Identifiers  map[string]Value
	Label        string
	Instructions []Value
	Predecessors []*Block
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
	successor.Predecessors = append(successor.Predecessors, block)
}

// Append adds a new value to the block.
func (block *Block) Append(value Value) {
	block.Instructions = append(block.Instructions, value)
}

// CanReachPredecessor checks if the `other` block appears as a predecessor or is the block itself.
func (block *Block) CanReachPredecessor(other *Block) bool {
	return block.CanReachPredecessor2(other, make(map[*Block]bool))
}

// CanReachPredecessor2 checks if the `other` block appears as a predecessor or is the block itself.
func (block *Block) CanReachPredecessor2(other *Block, traversed map[*Block]bool) bool {
	if other == block {
		return true
	}

	if traversed[block] {
		return false
	}

	traversed[block] = true

	for _, pre := range block.Predecessors {
		if pre.CanReachPredecessor2(other, traversed) {
			return true
		}
	}

	return false
}

// Contains checks if the value exists within the block.
func (block *Block) Contains(value Value) bool {
	return block.Index(value) != -1
}

// FindIdentifier searches for all the possible values the identifier
// can have and combines them to a phi instruction if necessary.
func (block *Block) FindIdentifier(name string) (value Value, exists bool) {
	return block.findIdentifier(name, make(map[*Block]Value))
}

// findIdentifier searches for all the possible values the identifier
// can have and combines them to a phi instruction if necessary.
func (block *Block) findIdentifier(name string, traversed map[*Block]Value) (Value, bool) {
	cached, isTraversed := traversed[block]

	if isTraversed {
		return cached, cached != nil
	}

	value, exists := block.Identifiers[name]

	if exists {
		traversed[block] = value
		return value, true
	}

	traversed[block] = nil

	switch len(block.Predecessors) {
	case 0:
		return nil, false
	case 1:
		value, exists := block.Predecessors[0].findIdentifier(name, traversed)

		if exists {
			traversed[block] = value
		}

		return value, exists
	default:
		values := set.Ordered[Value]{}

		for _, pre := range block.Predecessors {
			value, exists := pre.findIdentifier(name, traversed)

			if !exists {
				return nil, false
			}

			values.Add(value)
		}

		if values.Count() == 0 {
			return nil, false
		}

		if values.Count() == 1 {
			traversed[block] = values.Slice()[0]
			return values.Slice()[0], true
		}

		phi := &Phi{Arguments: values.Slice(), Typ: values.Slice()[0].Type()}
		block.Append(phi)
		block.Identify(name, phi)
		traversed[block] = phi
		return phi, true
	}
}

// Identify adds a new identifier or changes an existing one.
func (block *Block) Identify(name string, value Value) {
	if block.Identifiers == nil {
		block.Identifiers = make(map[string]Value, 8)
	}

	block.Identifiers[name] = value
}

// Index returns the position of the value or -1 if it doesn't exist within the block.
func (block *Block) Index(search Value) int {
	for i, value := range block.Instructions {
		if value == search {
			return i
		}
	}

	return -1
}

// InsertAt inserts the `value` at the given `index`.
func (block *Block) InsertAt(value Value, index int) {
	block.Instructions = slices.Insert(block.Instructions, index, value)
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