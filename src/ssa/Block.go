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
func (b *Block) AddSuccessor(successor *Block) {
	successor.Predecessors = append(successor.Predecessors, b)
}

// Append adds a new value to the block.
func (b *Block) Append(value Value) {
	b.Instructions = append(b.Instructions, value)
}

// CanReachPredecessor checks if the `other` block appears as a predecessor or is the block itself.
func (b *Block) CanReachPredecessor(other *Block) bool {
	return b.canReachPredecessor(other, make(map[*Block]bool))
}

// canReachPredecessor checks if the `other` block appears as a predecessor or is the block itself.
func (b *Block) canReachPredecessor(other *Block, traversed map[*Block]bool) bool {
	if other == b {
		return true
	}

	if traversed[b] {
		return false
	}

	traversed[b] = true

	for _, pre := range b.Predecessors {
		if pre.canReachPredecessor(other, traversed) {
			return true
		}
	}

	return false
}

// Contains checks if the value exists within the block.
func (b *Block) Contains(value Value) bool {
	return b.Index(value) != -1
}

// FindExisting returns an equal instruction that's already appended or `nil` if none could be found.
func (b *Block) FindExisting(instr Value) Value {
	if !instr.IsConst() {
		return nil
	}

	for _, existing := range slices.Backward(b.Instructions) {
		if existing.IsConst() && instr.Equals(existing) {
			return existing
		}

		// If we encounter a call, we can't be sure that the value is still the same.
		// TODO: This is a bit too conservative. We could check if the call affects the value.
		switch existing.(type) {
		case *Call, *CallExtern:
			return nil
		}
	}

	return nil
}

// FindIdentifier searches for all the possible values the identifier
// can have and combines them to a phi instruction if necessary.
func (b *Block) FindIdentifier(name string) (value Value, exists bool) {
	return b.findIdentifier(name, make(map[*Block]Value))
}

// findIdentifier searches for all the possible values the identifier
// can have and combines them to a phi instruction if necessary.
func (b *Block) findIdentifier(name string, traversed map[*Block]Value) (Value, bool) {
	if cached, isTraversed := traversed[b]; isTraversed {
		return cached, cached != nil
	}

	if value, exists := b.Identifiers[name]; exists {
		traversed[b] = value
		return value, true
	}

	traversed[b] = nil

	switch len(b.Predecessors) {
	case 0:
		return nil, false
	case 1:
		value, exists := b.Predecessors[0].findIdentifier(name, traversed)

		if exists {
			traversed[b] = value
		}

		return value, exists
	default:
		var values []Value

		for _, pre := range b.Predecessors {
			value, exists := pre.findIdentifier(name, traversed)

			if !exists {
				return nil, false
			}

			values = append(values, value)
			traversed[b] = value
		}

		if len(values) == 0 {
			return nil, false
		}

		if allSame(values) {
			return values[0], true
		}

		phi := &Phi{Arguments: values, Typ: values[0].Type()}
		b.InsertAt(phi, 0)
		b.Identify(name, phi)
		traversed[b] = phi
		return phi, true
	}
}

// IdentifiersFor returns an iterator for all the identifiers pointing to the given value.
func (b *Block) IdentifiersFor(value Value) iter.Seq[string] {
	return func(yield func(string) bool) {
		for name, val := range b.Identifiers {
			if val == value {
				if !yield(name) {
					return
				}
			}
		}
	}
}

// Identify adds a new identifier or changes an existing one.
func (b *Block) Identify(name string, value Value) {
	if b.Identifiers == nil {
		b.Identifiers = make(map[string]Value, 8)
	}

	b.Identifiers[name] = value
}

// IsIdentified returns true if the value can be obtained from one of the identifiers.
func (b *Block) IsIdentified(value Value) bool {
	for _, existing := range b.Identifiers {
		if existing == value {
			return true
		}
	}

	return false
}

// Index returns the position of the value or -1 if it doesn't exist within the block.
func (b *Block) Index(search Value) int {
	for i, value := range b.Instructions {
		if value == search {
			return i
		}
	}

	return -1
}

// InsertAt inserts the `value` at the given `index`.
func (b *Block) InsertAt(value Value, index int) {
	b.Instructions = slices.Insert(b.Instructions, index, value)
}

// Last returns the last value.
func (b *Block) Last() Value {
	return b.Instructions[len(b.Instructions)-1]
}

// RemoveAt sets the value at the given index to nil.
func (b *Block) RemoveAt(index int) {
	value := b.Instructions[index]

	for _, input := range value.Inputs() {
		input.RemoveUser(value)
	}

	b.Instructions[index] = nil
}

// RemoveNilValues removes all nil values from the block.
func (b *Block) RemoveNilValues() {
	b.Instructions = slices.DeleteFunc(b.Instructions, func(value Value) bool {
		return value == nil
	})
}

// ReplaceAllUses replaces all uses of `old` with `new`.
func (b *Block) ReplaceAllUses(old Value, new Value) {
	for _, instr := range b.Instructions {
		instr.Replace(old, new)
	}
}

// String returns the block label.
func (b *Block) String() string {
	return b.Label
}