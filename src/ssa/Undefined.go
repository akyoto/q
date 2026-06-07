package ssa

// Undefined is used in phi values to represent an undefined value.
var Undefined = &undefined{}

// undefined is used as a nil replacement.
type undefined struct {
	Independent
	Void
}

// Equals always returns false.
func (v *undefined) Equals(Value) bool { return false }

// String returns a human-readable representation of the undefined value.
func (v *undefined) String() string { return "undefined" }