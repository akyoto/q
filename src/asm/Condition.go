package asm

// Condition represents the exact condition to use in jumps
// and conditional set instructions. It separates signed and
// unsigned comparisons.
type Condition uint8

const (
	None Condition = iota
	Equal
	NotEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	UnsignedGreater
	UnsignedGreaterEqual
	UnsignedLess
	UnsignedLessEqual
)