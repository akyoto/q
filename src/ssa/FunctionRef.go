package ssa

// FunctionRef is separate from Function and defines
// the interface for the stored function reference.
type FunctionRef interface {
	Name() string
	Package() string
}