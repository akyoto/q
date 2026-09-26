package ssa

import (
	"iter"
	"slices"

	"git.urbach.dev/cli/q/src/set"
)

// Identifiers contains the identifier mappings before and after a block executed.
type Identifiers struct {
	Before set.CowMap[string, Value]
	After  set.CowMap[string, Value]
}

// FindIdentifier looks up an identifier.
func (i *Identifiers) FindIdentifier(name string) (Value, bool) {
	return i.After.Get(name)
}

// IdentifiersFor returns an iterator for all the identifiers pointing at the given value.
func (i *Identifiers) IdentifiersFor(value Value) iter.Seq[string] {
	return func(yield func(string) bool) {
		names := make([]string, 0, i.After.Count())

		for name, existing := range i.After.Raw() {
			if existing == value {
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
	i.After.Set(name, value)
}

// IsIdentified returns true if the value can be obtained from one of the identifiers.
func (i *Identifiers) IsIdentified(value Value) bool {
	for _, existing := range i.After.Raw() {
		if existing == value {
			return true
		}
	}

	return false
}

// ReplaceIdentifier replaces an existing identifier.
func (i *Identifiers) ReplaceIdentifier(name string, oldValue Value, newValue Value) {
	i.Before.Set(name, newValue)
	existing, _ := i.After.Get(name)

	if existing == oldValue {
		i.After.Set(name, newValue)
	}
}

// Unidentify deletes the identifier for the given value.
func (i *Identifiers) Unidentify(value Value) {
	i.After.RemoveByValue(value)
}