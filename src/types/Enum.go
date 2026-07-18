package types

// Enum is a namespaced set of constants.
type Enum struct {
	pkg     string
	name    string
	file    any
	members map[string]any
}

// NewEnum creates a new enum type.
func NewEnum(pkg string, name string, file any) *Enum {
	return &Enum{
		pkg:     pkg,
		name:    name,
		file:    file,
		members: make(map[string]any),
	}
}

// AddMember adds a named member to the enum.
func (e *Enum) AddMember(name string, value any) {
	e.members[name] = value
}

// File returns the file where the enum was defined.
func (e *Enum) File() any {
	return e.file
}

// Member returns the value for the given member name, or nil if not found.
func (e *Enum) Member(name string) (any, bool) {
	v, ok := e.members[name]
	return v, ok
}

// Name returns the name of the enum.
func (e *Enum) Name() string {
	return e.name
}

// Package returns the package of the enum.
func (e *Enum) Package() string {
	return e.pkg
}

// Size returns the size of the enum (same as int).
func (e *Enum) Size() int {
	return Int.Size()
}