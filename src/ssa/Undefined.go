package ssa

// Undefined is used in phi values to represent an undefined value.
var Undefined = &undefined{}

type undefined struct{ Void }

// Equals always returns false.
func (v *undefined) Equals(Value) bool { return false }

// Inputs always returns nil.
func (v *undefined) Inputs() []Value { return nil }

// Replace does nothing.
func (v *undefined) Replace(Value, Value) {}

// String returns a human-readable representation of the undefined value.
func (v *undefined) String() string { return "undefined" }