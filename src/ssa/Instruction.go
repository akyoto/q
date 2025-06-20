package ssa

// Instruction is a "fat struct" for performance reasons.
// It contains all the fields necessary to represent all instruction types.
type Instruction struct {
	Type       Type
	Parameters []Index
	Int        int64
	Float      float64
	String     string
}