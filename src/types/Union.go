package types

import "strings"

// Union is a sum type where any of the specified types can be the real type.
type Union struct {
	Types []Type
}

// Index returns the position of the given type within the union.
func (u *Union) Index(search Type) int {
	for i, typ := range u.Types {
		if Is(search, typ) {
			return i
		}
	}

	return -1
}

// Name returns the type name.
func (u *Union) Name() string {
	b := strings.Builder{}

	for i, typ := range u.Types {
		b.WriteString(typ.Name())

		if i != len(u.Types)-1 {
			b.WriteString(" | ")
		}
	}

	return b.String()
}

// Size returns the total size in bytes.
func (u *Union) Size() int {
	size := 0

	for _, typ := range u.Types {
		size = max(size, typ.Size())
	}

	return size
}

// String returns the type name.
func (u *Union) String() string {
	return u.Name()
}