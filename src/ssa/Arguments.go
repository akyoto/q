package ssa

import (
	"fmt"
	"slices"
	"strings"
)

// Arguments defines a list of values that this value depends on.
type Arguments []Value

// allSame checks if all elements are the same.
func allSame[T comparable](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	first := slice[0]

	for _, v := range slice[1:] {
		if v != first {
			return false
		}
	}

	return true
}

// Equals returns true if all the arguments are equal.
func (a Arguments) Equals(b Arguments) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Index returns the position of the value within the slice.
func (v Arguments) Index(search Value) int {
	return slices.Index(v, search)
}

// Inputs is the arguments list itself.
func (v Arguments) Inputs() []Value {
	return v
}

// Replace replaces the `old` value with `new`.
func (v Arguments) Replace(old Value, new Value) {
	for i, arg := range v {
		if arg == old {
			v[i] = new
		}
	}
}

// String returns a comma-separated list of all arguments.
func (v Arguments) String() string {
	tmp := strings.Builder{}

	for i, arg := range v {
		tmp.WriteString(fmt.Sprintf("%p", arg))

		if i != len(v)-1 {
			tmp.WriteString(", ")
		}
	}

	return tmp.String()
}