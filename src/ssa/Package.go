package ssa

// Package is a package reference.
type Package struct {
	Name     string
	IsExtern bool
	Independent
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

// String returns the name of the package.
func (p *Package) String() string { return p.Name }