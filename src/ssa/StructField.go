package ssa

// StructField is an interface for values that can be part of a struct.
type StructField interface {
	// Struct returns the struct that this field is a part of.
	Struct() *Struct
}