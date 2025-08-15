package types

import "git.urbach.dev/cli/q/src/token"

// Field is a memory region in a data structure.
type Field struct {
	Type     Type
	Name     string
	Tokens   token.List
	Position token.Position
	Index    uint8
	Offset   uint8
}

// String returns the name of the struct.
func (f *Field) String() string {
	return f.Name
}