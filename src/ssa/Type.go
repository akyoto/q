package ssa

// Type represents the instruction type.
type Type byte

const (
	None Type = iota

	// Values
	Int
	Float
	Func
	Register
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

	// Control flow
	If
	Jump
	Call
	Return
	Syscall

	// Special
	Phi
)