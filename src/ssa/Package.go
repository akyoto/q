package ssa

// Package is a package reference.
type Package struct {
	Name     string
	IsExtern bool
	Void
}

func (v *Package) Inputs() []Value      { return nil }
func (v *Package) Replace(Value, Value) {}
func (v *Package) String() string       { return v.Name }

func (a *Package) Equals(v Value) bool {
	b, sameType := v.(*Package)

	if !sameType {
		return false
	}

	return a.Name == b.Name
}