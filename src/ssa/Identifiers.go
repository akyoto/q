package ssa

import (
	"iter"
	"slices"
)

// Identifiers contains the identifier mappings before and after a block executed.
type Identifiers struct {
	Before map[string]Value
	After  map[string]Value
}

// FindIdentifier looks up an identifier.
func (i *Identifiers) FindIdentifier(name string) (Value, bool) {
	value, exists := i.After[name]
	return value, exists
}

// IdentifiersFor returns an iterator for all the identifiers pointing to the given value.
func (i *Identifiers) IdentifiersFor(value Value) iter.Seq[string] {
	return func(yield func(string) bool) {
		names := make([]string, 0, len(i.After))

		for name, val := range i.After {
			if val == value {
				names = append(names, name)
			}
		}

		slices.Sort(names)

		for _, name := range names {
			if !yield(name) {
				return
			}
		}
	}
}

// Identify adds a new identifier or changes an existing one.
func (i *Identifiers) Identify(name string, value Value) {
	if i.After == nil {
		i.After = make(map[string]Value, 8)
	}

	i.After[name] = value
}

// IsIdentified returns true if the value can be obtained from one of the identifiers.
func (i *Identifiers) IsIdentified(value Value) bool {
	for _, existing := range i.After {
		if existing == value {
			return true
		}
	}

	return false
}

// ReplaceIdentifier replaces an existing identifier.
func (i *Identifiers) ReplaceIdentifier(name string, oldValue Value, newValue Value) {
	i.Before[name] = newValue

	if i.After[name] == oldValue {
		i.After[name] = newValue
	}
}

// Unidentify deletes the identifier for the given value.
func (i *Identifiers) Unidentify(value Value) {
	for name, existing := range i.After {
		if existing == value {
			delete(i.After, name)
			return
		}
	}
}