package ssa

// Package is a package reference.
type Package struct {
	Name     string
	IsExtern bool
	Void
}

// Equals returns true if the packages are equal.
func (a *Package) Equals(v Value) bool {
	b, sameType := v.(*Package)

	if !sameType {
		return false
	}

	return a.Name == b.Name
}

// Inputs returns nil because a package has no inputs.
func (p *Package) Inputs() []Value { return nil }

// Replace does nothing because a package has no inputs.
func (p *Package) Replace(Value, Value) {}

// String returns the name of the package.
func (p *Package) String() string { return p.Name }