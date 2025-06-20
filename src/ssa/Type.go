package ssa

// Type represents the instruction type.
type Type byte

const (
	None Type = iota

	// Values
	Int
	Float
	String

	// Binary
	Add
	Sub
	Mul
	Div
	Mod

	// Bitwise
	And
	Or
	Xor
	Shl
	Shr

	// Branch
	If
	Jump

	// Special
	Call
	Phi
)