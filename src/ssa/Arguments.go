package ssa

import (
	"strings"
)

// Arguments defines a list of values that this value depends on.
type Arguments []Value

// Equals returns true if all the arguments are equal.
func (a Arguments) Equals(b Arguments) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
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
		tmp.WriteString(arg.String())

		if i != len(v)-1 {
			tmp.WriteString(", ")
		}
	}

	return tmp.String()
}