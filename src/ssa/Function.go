package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// Function is a reference to a function.
type Function struct {
	Typ *types.Function
	FunctionRef
	Independent
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

// IsPure returns true because a function reference is always constant.
func (f *Function) IsPure() bool { return true }

// String returns a human-readable representation of the function.
func (f *Function) String() string {
	if f.Package() == "" {
		return f.Name()
	}

	return f.Package() + "." + f.Name()
}

// Type returns the type of the function.
func (f *Function) Type() types.Type { return f.Typ }