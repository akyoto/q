package ssa

// Undefined is used in phi values to represent an undefined value.
var Undefined = &undefined{}

// undefined is used as a nil replacement.
type undefined struct {
	Independent
	Void
}

// Equals only returns true when the value points to the same object.
func (a *undefined) Equals(b Value) bool { return a == b }

// String returns a human-readable representation of the undefined value.
func (v *undefined) String() string { return "undefined" }