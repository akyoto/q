package ssa

import (
	"iter"
	"slices"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Identifiers  map[string]Value
	Loop         *Block
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
	return block.canReachPredecessor(other, make(map[*Block]bool))
}

// canReachPredecessor checks if the `other` block appears as a predecessor or is the block itself.
func (block *Block) canReachPredecessor(other *Block, traversed map[*Block]bool) bool {
	if other == block {
		return true
	}

	if traversed[block] {
		return false
	}

	traversed[block] = true

	for _, pre := range block.Predecessors {
		if pre.canReachPredecessor(other, traversed) {
			return true
		}
	}

	return false
}

// Contains checks if the value exists within the block.
func (block *Block) Contains(value Value) bool {
	return block.Index(value) != -1
}

// FindExisting returns an equal instruction that's already appended or `nil` if none could be found.
func (block *Block) FindExisting(instr Value) Value {
	if !instr.IsConst() {
		return nil
	}

	for _, existing := range block.Instructions {
		if existing.IsConst() && instr.Equals(existing) {
			return existing
		}
	}

	return nil
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
		var values []Value

		for _, pre := range block.Predecessors {
			value, exists := pre.findIdentifier(name, traversed)

			if !exists {
				return nil, false
			}

			values = append(values, value)
			traversed[block] = value
		}

		if len(values) == 0 {
			return nil, false
		}

		if allSame(values) {
			return values[0], true
		}

		phi := &Phi{Arguments: values, Typ: values[0].Type()}
		block.InsertAt(phi, 0)
		block.Identify(name, phi)
		traversed[block] = phi
		return phi, true
	}
}

// IdentifiersFor returns an iterator for all the identifiers pointing to the given value.
func (block *Block) IdentifiersFor(value Value) iter.Seq[string] {
	return func(yield func(string) bool) {
		for name, val := range block.Identifiers {
			if val == value {
				if !yield(name) {
					return
				}
			}
		}
	}
}

// Identify adds a new identifier or changes an existing one.
func (block *Block) Identify(name string, value Value) {
	if block.Identifiers == nil {
		block.Identifiers = make(map[string]Value, 8)
	}

	block.Identifiers[name] = value
}

// IsIdentified returns true if the value can be obtained from one of the identifiers.
func (block *Block) IsIdentified(value Value) bool {
	for _, existing := range block.Identifiers {
		if existing == value {
			return true
		}
	}

	return false
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

// Last returns the last value.
func (block *Block) Last() Value {
	return block.Instructions[len(block.Instructions)-1]
}

// RemoveNilValues removes all nil values from the block.
func (block *Block) RemoveNilValues() {
	block.Instructions = slices.DeleteFunc(block.Instructions, func(value Value) bool {
		return value == nil
	})
}

// ReplaceAllUses replaces all uses of `old` with `new`.
func (block *Block) ReplaceAllUses(old Value, new Value) {
	for _, instr := range block.Instructions {
		instr.Replace(old, new)
	}
}

// String returns the block label.
func (block *Block) String() string {
	return block.Label
}