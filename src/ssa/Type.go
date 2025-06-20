package ssa

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

	// Special
	Call
	Phi
)