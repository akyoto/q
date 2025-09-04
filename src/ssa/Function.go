package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// Function is a reference to a function.
type Function struct {
	Typ *types.Function
	FunctionRef
	Liveness
	Source
}

// Equals returns true if the functions are equal.
func (a *Function) Equals(v Value) bool {
	b, sameType := v.(*Function)

	if !sameType {
		return false
	}

	return a.Package() == b.Package() && a.Name() == b.Name()
}

// Inputs returns nil because a function reference has no inputs.
func (f *Function) Inputs() []Value { return nil }

// IsConst returns true because a function reference is always constant.
func (f *Function) IsConst() bool { return true }

// Replace does nothing because a function reference has no inputs.
func (f *Function) Replace(Value, Value) {}

// String returns a human-readable representation of the function.
func (f *Function) String() string {
	if f.Package() == "" {
		return f.Name()
	}

	return f.Package() + "." + f.Name()
}

// Type returns the type of the function.
func (f *Function) Type() types.Type { return f.Typ }