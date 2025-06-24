package set

import (
	"iter"
	"slices"
)

// Ordered is an ordered set.
type Ordered[T comparable] struct {
	values []T
}

// Add adds a value to the set if it doesn't exist yet.
// It returns `false` if it already exists, `true` if it was added.
func (set *Ordered[T]) Add(value T) bool {
	if slices.Contains(set.values, value) {
		return false
	}

	set.values = append(set.values, value)
	return true
}

// All returns an iterator over all the values in the set.
func (set *Ordered[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range set.values {
			if !yield(value) {
				return
			}
		}
	}
}

// Count returns the number of elements in the set.
func (set *Ordered[T]) Count() int {
	return len(set.values)
}