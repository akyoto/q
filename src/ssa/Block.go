package ssa

import (
	"maps"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Identifiers
	Protected    map[Value][]Value
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
	if slices.Contains(successor.Predecessors, b) {
		return
	}

	successor.Predecessors = append(successor.Predecessors, b)

	if len(b.Protected) > 0 {
		if successor.Protected == nil {
			successor.Protected = make(map[Value][]Value, len(b.Protected))
		}

		maps.Copy(successor.Protected, b.Protected)
	}

	if b.Identifiers.After == nil {
		return
	}

	if successor.Identifiers.After == nil {
		successor.Identifiers.Before = make(map[string]Value, len(b.Identifiers.After))
		successor.Identifiers.After = make(map[string]Value, len(b.Identifiers.After))

		if len(successor.Predecessors) == 1 {
			maps.Copy(successor.Identifiers.Before, b.Identifiers.After)
			maps.Copy(successor.Identifiers.After, b.Identifiers.After)
			return
		}
	}

	keys := make([]string, 0, max(len(b.Identifiers.After), len(successor.Identifiers.After)))

	for name := range successor.Identifiers.Before {
		if !slices.Contains(keys, name) {
			keys = append(keys, name)
		}
	}

	for name := range b.Identifiers.After {
		if !slices.Contains(keys, name) {
			keys = append(keys, name)
		}
	}

	slices.SortFunc(keys, func(a string, b string) int {
		return strings.Compare(b, a)
	})

	var modifiedStructs []string

	for _, name := range keys {
		oldValue, oldExists := successor.Identifiers.Before[name]
		newValue, newExists := b.Identifiers.After[name]

		switch {
		case oldExists:
			if oldValue == newValue {
				continue
			}

			_, isStruct := oldValue.(*Struct)

			if isStruct {
				modifiedStructs = append(modifiedStructs, name)
				continue
			}

			definedLocally := successor.Index(oldValue) != -1

			if definedLocally {
				phi, isPhi := oldValue.(*Phi)

				if isPhi {
					if newExists {
						phi.Arguments = append(phi.Arguments, newValue)
					} else {
						phi.Arguments = append(phi.Arguments, Undefined)
					}
				}

				continue
			}

			phi := &Phi{
				Arguments: make([]Value, len(successor.Predecessors)-1, len(successor.Predecessors)),
				Typ:       oldValue.Type(),
			}

			for i := range phi.Arguments {
				phi.Arguments[i] = oldValue
			}

			successor.InsertAt(0, phi)
			successor.ReplaceIdentifier(name, oldValue, phi)

			if newExists {
				phi.Arguments = append(phi.Arguments, newValue)
			} else {
				phi.Arguments = append(phi.Arguments, Undefined)
			}

		case newExists:
			phi := &Phi{
				Arguments: make([]Value, len(successor.Predecessors)-1, len(successor.Predecessors)),
				Typ:       newValue.Type(),
			}

			for i := range phi.Arguments {
				phi.Arguments[i] = Undefined
			}

			successor.InsertAt(0, phi)
			successor.ReplaceIdentifier(name, oldValue, phi)
			phi.Arguments = append(phi.Arguments, newValue)
		}
	}

	// Structs that were modified in branches need to be recreated
	// to use the new Phi values as their arguments.
	for _, name := range modifiedStructs {
		structure := successor.Identifiers.Before[name].(*Struct)
		structType := types.Unwrap(structure.Typ).(*types.Struct)
		newStruct := &Struct{Typ: structure.Typ, Arguments: make(Arguments, len(structure.Arguments))}

		for i, field := range structType.Fields {
			newStruct.Arguments[i] = successor.Identifiers.Before[name+"."+field.Name]
		}

		successor.ReplaceIdentifier(name, structure, newStruct)
	}
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
	if !instr.IsPure() {
		return nil
	}

	for _, existing := range slices.Backward(b.Instructions) {
		if existing.IsPure() && instr.Equals(existing) {
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
func (b *Block) InsertAt(index int, values ...Value) {
	b.Instructions = slices.Insert(b.Instructions, index, values...)
}

// Last returns the last value.
func (b *Block) Last() Value {
	if len(b.Instructions) == 0 {
		return nil
	}

	return b.Instructions[len(b.Instructions)-1]
}

// Phis is an iterator for all phis at the top of the block.
func (b *Block) Phis(yield func(*Phi) bool) {
	for _, instr := range b.Instructions {
		phi, isPhi := instr.(*Phi)

		if !isPhi || !yield(phi) {
			return
		}
	}
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
	return CleanLabel(b.Label)
}

// Protect protects the given value from being accessed before the error value is checked.
func (b *Block) Protect(err Value, protected []Value) {
	if b.Protected == nil {
		b.Protected = make(map[Value][]Value)
	}

	b.Protected[err] = protected
}

// Unprotect stops protecting the variables for the given error value.
func (b *Block) Unprotect(err Value) {
	delete(b.Protected, err)
}